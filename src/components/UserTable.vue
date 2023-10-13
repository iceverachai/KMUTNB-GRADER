<script setup lang="ts">
import TableComponent from '../components/TableComponent.vue'
import { reactive, onMounted } from 'vue';
import { RouterLink } from 'vue-router';
import EditIcon from "./EditIcon.vue";

defineEmits(["onUpdateUserClick"])

const sortState = reactive<{ col: string }>({ col: '' })

const prop = defineProps<{tabledata:TableData}>()

const data = reactive<TableData>(
    {} as TableData
)


const SortMethod = (SortBy: string) => {
    sortState.col = SortBy;
    if (SortBy === "Assigned") {
        const keyof = Object.keys(data.response.Structure)
        if (!data.response.Assigned || !keyof.length) return;
        data.response.Assigned.sort((a: any, b: any) => {
            return b.Score - a.Score;
        });
    } else if(SortBy === "Namesort"){
        const keyof = Object.keys(data.response.Structure)
        if (!data.response.Assigned || !keyof.length) return;
        data.response.Assigned.sort((a: any, b: any) => {
            return ('' + a.Firstname).localeCompare(b.Firstname)
        });
    }else {
        const keyof = Object.keys(data.response.Structure)
        if (!data.response.Assigned || !keyof.length) return;
        data.response.Assigned.sort((a: any, b: any) => {
            const compareA = a.Data[SortBy] ? a.Data[SortBy].Score : -1
            const compareB = b.Data[SortBy] ? b.Data[SortBy].Score : -1
            
            return compareB - compareA;
        });

    }
}

const getProgressStyle = (p: Assignment) => {
    const keyof = Object.keys(data.response.Structure)
    if (!p || !keyof.length) return;
    let count = 0;
    keyof.forEach((data: string) => {
        if (p[data] && p[data].Pass) {
            count++;
        }
    })
    const width = (count / keyof.length) * 100;
    
    let color;

    if(width >= 100){
        color = "rgb(35, 197, 35)"
    }else if(width >= 50){
        color = "rgb(255, 162, 0)"
    }else{
        color = "rgb(246, 49, 49)"
    }

    return [{ c: count, style: `width:${width}%; background-color: ${color};` }]
}


const UserMapping = (UserData: any): string => {
    if (!UserData) return "unknown";

    return UserData.Pass ? 'checked' : 'cross';
}

onMounted(async() => {
    data.response = prop.tabledata.response
    data.status = prop.tabledata.status
    if(data.status === "success")
        SortMethod('Namesort');
})
</script>

<template>
    <table class="Table" v-if="data.status === 'success'">
        <thead>
            <tr>
                <th style="z-index:15" class="UserHeader" @click="SortMethod('Namesort')">ชื่อ-สกุล <label v-if="'Namesort' === sortState.col">&#9660;</label></th>
                <!-- <th style="z-index:10" class="UserHeader" >สกุล</th> -->
                <th style="z-index:14" @click="SortMethod('Assigned')">{{ Object.keys(data.response.Structure).length }} ข้อ

                    <label v-if="'Assigned' === sortState.col">&#9660;</label>
                </th>
                <th class="UserHeader" style="text-align:center;" v-for="(Propo, index) in data.response.Structure" @click="SortMethod(Propo.Id)"
                    :key="index">
                    <div>
                        <span>{{ Propo.Name }} </span>
                        <label style="color: rgb(14, 192, 14);">{{ " ( " + Propo.Maxscore + " )" }}</label>
                        <label v-if="Propo.Id === sortState.col">&#9660;</label> 
                    </div>
                </th>
            </tr>
        </thead>
        
        <TransitionGroup tag="tbody" name="list">
            <tr v-for="(UserData, index) in data?.response?.Assigned" :key="UserData.Username+index">
                <td class="UserHeader" style="z-index:10">
                    {{ UserData.Firstname + " " }}
                    {{ UserData.Lastname }}
                    <EditIcon @click="$emit('onUpdateUserClick',UserData.Username)"/>
                </td>

                <th>
                    <div class="progbar"
                        v-for="(item) in getProgressStyle(UserData.Data)">
                        <div class="inner-progbar" :style="item.style"></div>
                        <div class="numcount">
                            {{ item.c }} / {{ Object.keys(data.response.Structure).length }}
                        </div>
                    </div>
                </th>
                <!-- แก้เป็น ชื้อ-นามสกุล -->
                <td class="submitresult" v-for="(assignment, _index2) in data?.response?.Structure" :key="assignment.Id + UserData.Username">
                    <div class="attempt" v-if="UserData.Data[assignment.Id]">{{  UserData.Data[assignment.Id].Attempt  }}</div>
                    <div class="score" v-if="UserData.Data[assignment.Id] && !UserData.Data[assignment.Id].Pass">{{  UserData.Data[assignment.Id].Score  }}</div>
                    <router-link :to="`/editor/code-${assignment?.Id}/${UserData.Username}/${(UserData.Firstname)}/${(UserData.Lastname)}`"
                        v-if="UserData.Data[assignment.Id]!==undefined" target='_blank'>
                        <TableComponent :type="UserMapping(UserData.Data[assignment.Id])">ยังไม่ส่ง
                        </TableComponent>
                    </router-link>
                    <TableComponent :type="UserMapping(UserData.Data[assignment.Id])" v-else>ยังไม่ส่ง
                    </TableComponent>
                    <div class="timestamp" v-if="UserData.Data[assignment.Id]">{{ new Date(UserData.Data[assignment.Id].Timestamp * 1000).toLocaleString() }}</div>
                </td>
            </tr>
        </TransitionGroup>

        
    </table>
</template>

<style scoped>
.progbar{
    width:7rem;
    height:1rem;
    background-color: rgb(177, 177, 177);
    z-index:-1;
    padding: 0.8rem 0;
    position: relative;
}
.inner-progbar{
    position: absolute;
    top: 0;
    left: 0;
    height: 100%;
    z-index:-1;
}
.numcount{
    position: absolute;
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #fff;
    top: 0;
    font-size: 1rem;
}
.list-move, /* apply transition to moving elements */
.list-enter-active,
.list-leave-active {
  transition: all 0.2s ease-in-out;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
}

.list-leave-active {
  position: absolute;
}
.submitresult{
    position: relative;
}
.timestamp{
    position: absolute;
    right: 50%;
    transform: translate(50%, 0);
    bottom: 0;
    color: #fff;
    font-size: 0.8rem;
}
.score{
    position: absolute;
    top: 0;
    right: 0;
    color: #fff;
    font-size: 0.8rem;
}
.attempt{
    position: absolute;
    top: 0;
    left: 5%;
    color: #fff;
    font-size: 0.8rem;
}
</style>