<script setup lang="ts">
import EditIcon from "./EditIcon.vue";
import { reactive, onMounted } from 'vue';
defineEmits(["onEditProbPopupBtnClick","TestcaseShowClick", "TemplateShowClick"])
type probStructure = {
    response: {
        [key: string]: {
            Id: string,
            Name: string,
            Info: string,
            Link: string,
            Testcase: any
        }
    },
    status: string
}
const data = reactive<probStructure>({} as probStructure)

const GetAssignmentTable = async () => {
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
            data.response = json.response;
            data.status = json.status;
            return
        } else {
            // data.status = json.status;
        }
    } catch (error) {
        data.status = "error";
    }
}

onMounted(() => {
    GetAssignmentTable()
})
</script>

<template>
    <table class="Table" v-if="data.status === 'success'">
        <thead>
            <tr style="text-align: center">
                <th style="z-index:10;text-align: center" class="UserHeader">ID</th>
                <th style="z-index:10;text-align: center" class="UserHeader">Name</th>
                <th style="z-index:10;text-align: center" class="UserHeader">Info</th>
                <th style="z-index:10;text-align: center" class="UserHeader">Link</th>
                <th style="z-index:10;text-align: center" class="UserHeader">Testcase</th>
                <th style="z-index:10;text-align: center" class="UserHeader">Template</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(Data, index) in data?.response" :key="index">
                <th class="" style="z-index:9;background-color:none; text-align: left;">{{ Data.Id }}
                    <EditIcon @click="$emit('onEditProbPopupBtnClick', Data)" />
                </th>
                <td style="z-index:10">{{ Data.Name }}</td>
                <td style="z-index:10"><a :href="Data.Info" target="_blank" rel="noopener noreferrer" v-if="Data.Info && Data.Info.length">Go to</a></td>
                <td style="z-index:10"><a :href="Data.Link" target="_blank" rel="noopener noreferrer">Go to</a></td>
                <td style="z-index:10;cursor:pointer;text-decoration:underline;" @click="$emit('TestcaseShowClick',Data)">Show</td>
                <td style="z-index:10;cursor:pointer;text-decoration:underline;" @click="$emit('TemplateShowClick',Data)">Show</td>
                <!-- {{ Data.Testcase }} -->
            </tr>
        </tbody>
    </table>
</template>

<style scoped>
a:link, a:visited{
    color: var(--ad-asm-link);
}

</style>