<script setup lang="ts">
import TableComponent from "../components/TableComponent.vue"
import EditIcon from "../components/EditIcon.vue"
import { reactive, onMounted } from 'vue';

type tabledata = { 
            Groupname: string,
            Member: {
                [key:string]: {
                    Username: string,
                    Firstname: string,
                    Lastname: string,
                    Score: number,
                    Work: {
                        [key:string]: {
                            Pass: boolean,
                            Timestamp: number
                            Score: number,
                            Attempt: number
                        }
                    }
                }
            },
            Assignment: {
                [key:string]: {
                    Id: string,
                    Name: string
                    Maxscore: number
                }
            }
}

const render = reactive<{order: Array<string>}>({order: []})

const sortState= reactive<{col : string}>({} as {col : string})

const props = defineProps<{tabledata: tabledata}>()

const SortMethod = (SortBy: string) => {
    sortState.col = SortBy;
    if (SortBy === "Assigned") {
        const keyof = Object.keys(props.tabledata.Assignment)
        const keymember = Object.keys(props.tabledata.Member)
        if (!props.tabledata.Member || !keyof.length) return;
        keymember.sort((a,b)=>{
            return getthisgroupScore(props.tabledata.Member[b].Work) - getthisgroupScore(props.tabledata.Member[a].Work);
        })
        render.order = keymember
    } else if(SortBy === "Namesort"){
        const keyof = Object.keys(props.tabledata.Member)
        const keymember = Object.keys(props.tabledata.Member)
        if (!props.tabledata.Member || !keyof.length) return;
        keymember.sort((a,b)=>{
            return ('' + props.tabledata.Member[a].Firstname).localeCompare(props.tabledata.Member[b].Firstname)
        })
        render.order = keymember
    }else {
        const keyof = Object.keys(props.tabledata.Assignment)
        const keymember = Object.keys(props.tabledata.Member)
        if (!props.tabledata.Member || !keyof.length) return;
        keymember.sort((a: any, b: any) => {
            const compareA = props.tabledata.Member[a].Work[SortBy] ? props.tabledata.Member[a].Work[SortBy].Score : -1
            const compareB = props.tabledata.Member[b].Work[SortBy] ? props.tabledata.Member[b].Work[SortBy].Score : -1
            
            return compareB - compareA;
        });
        render.order = keymember
    }
}

const getthisgroupScore = (p:any) => {
    const keyof = Object.keys(props.tabledata.Assignment);
    if (!p || !keyof.length) return 0;
    let count = 0;
    keyof.forEach((data: string) => {
        if (p[data] && p[data].Pass) {
            count++;
        }
    })
    return count;
}

const getProgressStyle = (p: any) => {
    const keyof = Object.keys(props.tabledata.Assignment)
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

const UserMapping = (Userkey: any): string => {
    if (!Userkey) return "unknown";

    return Userkey.Pass ? 'checked' : 'cross';
}

onMounted(() => { 
    render.order = Object.keys(props.tabledata.Member)
})

</script>

<template>
    <table class="Table">
        <thead>
            <tr>
                <th style="z-index:15" class="UserHeader" @click="SortMethod('Namesort')">ชื่อ-สกุล
                    <label v-if="'Namesort' === sortState.col">&#9660;</label>
                </th>
                <!-- <th style="z-index:10" class="UserHeader" >สกุล</th> -->
                <th style="z-index:14" @click="SortMethod('Assigned')">{{ Object.keys(props.tabledata.Assignment).length }} ข้อ

                    <label v-if="'Assigned' === sortState.col">&#9660;</label>
                </th>
                <th class="UserHeader" style="text-align:center;" v-for="(Propo, index) in props.tabledata.Assignment"
                    :key="index" @click="SortMethod(Propo.Id)">
                    <div>
                        <span>{{ Propo.Name }} </span>
                        <label style="color: rgb(14, 192, 14);">{{ " ( " + Propo.Maxscore + " )" }}</label>
                        <label v-if="Propo.Id === sortState.col">&#9660;</label> 
                    </div>
                </th>
            </tr>
        </thead>
        <TransitionGroup tag="tbody" name="list">
            <tr v-for="(Userkey, index) in render.order" :key="tabledata.Member[Userkey].Username+index">
                <td class="UserHeader" style="z-index:10">
                    {{ tabledata.Member[Userkey].Firstname + " " }}
                    {{ tabledata.Member[Userkey].Lastname }}
                    <EditIcon @click="$emit('onUpdateUserClick', tabledata.Member[Userkey].Username)" />
                </td>

                <th>
                    <div class="progbar"
                        v-for="(item) in getProgressStyle(tabledata.Member[Userkey].Work)">
                        <div class="inner-progbar" :style="item.style">
                        </div>
                        <div class="numcount">
                            {{ item.c }} / {{ Object.keys(props.tabledata.Assignment).length }}
                        </div>
                    </div>
                </th>
                <!-- แก้เป็น ชื้อ-นามสกุล -->
                <td class="submitresult" v-for="(assignment, _index2) in props.tabledata?.Assignment" :key="assignment.Id + tabledata.Member[Userkey].Username">
                    <div class="attempt" v-if="tabledata.Member[Userkey].Work[assignment.Id]">{{  tabledata.Member[Userkey].Work[assignment.Id].Attempt  }}</div>
                    <div class="score" v-if="tabledata.Member[Userkey].Work[assignment.Id] && !tabledata.Member[Userkey].Work[assignment.Id].Pass">{{ tabledata.Member[Userkey].Work[assignment.Id].Score }}</div>
                    <router-link
                        :to="`/editor/code-${assignment?.Id}/${tabledata.Member[Userkey].Username}/${(tabledata.Member[Userkey].Firstname)}/${(tabledata.Member[Userkey].Lastname)}`"
                        v-if="tabledata.Member[Userkey].Work[assignment.Id] !== undefined" target='_blank'>
                        <TableComponent :type="UserMapping(tabledata.Member[Userkey].Work[assignment.Id])">ยังไม่ส่ง
                        </TableComponent>
                    </router-link>
                    <TableComponent :type="UserMapping(tabledata.Member[Userkey].Work[assignment.Id])" v-else>ยังไม่ส่ง
                    </TableComponent>
                    <div class="timestamp" v-if="tabledata.Member[Userkey].Work[assignment.Id]">{{ new Date(tabledata.Member[Userkey].Work[assignment.Id].Timestamp * 1000).toLocaleString() }}</div>
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
    overflow: hidden;
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
    left: 0;
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