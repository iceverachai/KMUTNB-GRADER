<script setup lang="ts">
// import { useUserStore } from "../module/userstore";
import TerminateIcon from "./TerminateIcon.vue"
import { reactive, onMounted, } from 'vue';

const data = reactive<Groupdata>({} as Groupdata);
const Prob = reactive<GroupProbs>({} as GroupProbs)


const GetGrouptable = async () => {
    try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api + "/api/sentry/getgrouplist", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            }
        });
        const json = await response.json();
        if (json.status === "success" && json.response) {
            data.response = json.response;
            data.status = json.status;
            return
        } else {

        }
    } catch (error) {
        // data.status = "error";
    }
}

const GetProbtable = async () => {
    try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api + "/api/sentry/getassignmentlist", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            }
        });
        const json = await response.json();
        if (json.status === "success" && json.response) {
            Prob.response = json.response;
            Prob.status = json.status;
            
            return
        } else {

        }
    } catch (error) {
        // data.status = "error";
    }
}

const onMappingStateChange = async (event: any, index: any, index2: any) => {
    let payload;
    if (!event.target?.checked) {
        payload = {...data["response"][index2]["Assignment"]}
        delete payload[index];
    } else {
        payload = {...data["response"][index2]["Assignment"]}
        payload[index] = {
            id: index,
            sol_maxcase: false,
            sol_randomcase: false,
            sol_randomcasenum: 5,
            sol_randomcasenumrange: [4, 7]
        }
    } 
    if(!payload) return;
    try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api + "/api/sentry/updategroupassignment", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(
                {
                    groupname:data["response"][index2].Groupname,
                    assignment:payload
                }
            )
        });
        const json = await response.json();
        if (json.status === "success" && json.response) {
            data["response"][index2]["Assignment"]=payload;
        } else {
            console.log("onMappingStateChange Error")
        }
    } catch (error) {
        // data.status = "error";
    }
    
}

const onRemoveClickHandler = async (groupname:any) => {
    try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api + "/api/sentry/removegroup", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(
                {
                    groupname,
                }
            )
        });
        const json = await response.json();
        if (json.status === "success" && json.response) {

        } else {
            console.log("Updating Error", json)
        }
    } catch (error) {
        // data.status = "error";
    }
}

const onUpdatingGroupState = async (event: any, index: any) => {
    try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api + "/api/sentry/updategroupstate", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(
                {
                    groupname:data["response"][index].Groupname,
                    active:event.target?.checked
                }
            )
        });
        const json = await response.json();
        if (json.status === "success" && json.response) {
            data["response"][index]["Active"]=event.target?.checked;
        } else {
            console.log("onUpdatingGroupState Error")
        }
    } catch (error) {
        // data.status = "error";
    }
    
}

onMounted(() => {
    GetGrouptable()
    GetProbtable()
})
</script>

<template>
    <table class="Table" v-if="data.status === 'success'">
        <thead>
            <tr>
                <th style="z-index:11" class="UserHeader">โจทย์</th>
                <th class="UserHeader" style="z-index:10;text-align:center;" v-for="(renderdata, index) in data?.response"
                    :key="index">
                    <div class="controller">
                        <span>{{ renderdata.Groupname }}</span>
                        <input type="checkbox" name="groupstate" :checked="renderdata.Active"
                        @input="(e) => onUpdatingGroupState(e, index)" />
                        <TerminateIcon @click="onRemoveClickHandler(renderdata.Groupname)"/>
                    </div>

                </th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(Propo, index) in Prob.response" :key="index">
                <td class="UserHeader" style="z-index:10;">
                    {{ Propo.Name }}
                </td>
                <!-- แก้เป็น ชื้อ-นามสกุล -->
                <td v-for="(renderdata, index2) in data?.response" :key="index2">
                    <input type="checkbox" name="scales" :checked="renderdata.Assignment[index] ? true : false"
                        @input="(e) => onMappingStateChange(e, index, index2)" />
                    <!-- {{ renderdata.Assignment[Propo.Id] }} -->
                </td>
            </tr>
        </tbody>
    </table>
</template>

<style scoped>
.controller{
    display: flex;
    color: var(--ad-tableGroupF);
    justify-content: center;
    gap: 1rem;
}
.controller:deep() path{
    fill: var(--ad-tableGroupDelB);
}
</style>