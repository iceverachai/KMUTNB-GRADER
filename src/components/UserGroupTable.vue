<script setup lang="ts">
import { reactive, onMounted } from 'vue';
const userdata = reactive<UserData>({} as UserData)
const groupdata = reactive<Groupdata>({} as Groupdata);

const GetUserData = async () => {
    try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api + "/api/sentry/getuser", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            }
        });
        const json = await response.json();
        if (json.status === "success" && json.response) {
            userdata.response = json.response;
            userdata.status = json.status;
            return
        } else {
            userdata.status = json.status;
        }
    } catch (error) {
        userdata.status = "error";
    }
}

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
            groupdata.response = json.response;
            groupdata.status = json.status;
        }
    } catch (error) {
        // data.status = "error";
    }
}

const onChangeGroupHandler = async (e:any, username:string, pass_groups:Array<string>, groupname:string) => {
    try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        let groups = JSON.parse(JSON.stringify(pass_groups))

        if(e.target.checked === true){
            groups.push(groupname)
        }else{
            groups = groups.filter((e:string) => e !== groupname);
        }

        await fetch(import.meta.env.VITE_api + "/api/sentry/updateusergroup", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify({
                username,
                groups
            })
        });
    } catch (error) {
        // data.status = "error";
    }
}

onMounted(() => {
    GetUserData()
    GetGrouptable()
})
</script>

<template>
    <table class="Table" v-if="userdata.status === 'success'">
        <thead>
            <tr>
                <th style="z-index:10" class="UserHeader">User</th>
                <th style="z-index:10" class="UserHeader">ชื่อ-นามสกุล</th>
                <th v-for="(item, index) in groupdata?.response" :key="index">{{ item.Groupname }}</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(UserData, index) in userdata?.response" :key="index">
                <td class="UserHeader" style="z-index:10">{{ UserData.Username }}</td>
                <td class="UserHeader" style="z-index:10">{{ UserData.Firstname+" "+UserData.Lastname }}</td>
                <td v-for="(item, index2) in groupdata?.response" :key="index2" v-if="UserData.Groups">
                    <input type="checkbox" :checked="UserData.Groups.includes(item.Groupname)" @change="(e) => onChangeGroupHandler(e, UserData.Username, UserData.Groups, item.Groupname)" />
                </td>
            </tr>
        </tbody>
    </table>
</template>

<style scoped>


</style>