<script setup lang="ts">
import { reactive, onMounted } from 'vue';
defineEmits(["SelectedChange"])
const data = reactive<SelfAssignment>({} as SelfAssignment)

const group_data = reactive<GroupAssignment>({} as GroupAssignment);

const GetAssigned = async () => {
    try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api + "/api/getassignment", {
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
            data.status = json.status;
        }
    } catch (error) {
        data.status = "error";
    }
}

const GetGroupWork = async () => {
    try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api + "/api/getgroupwork", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            }
        });
        const json = await response.json();
        if (json.status === "success" && json.response) {
            group_data.response = json.response;
            group_data.status = json.status;
            return
        } else {
            group_data.status = json.status;
        }
    } catch (error) {
        data.status = "error";
    }
}

const SelfData = reactive<{
  response: {
    Firstname: string,
    Lastname: string,
    Groups: Array<string>,
    Work: {
        [key: string]: {
            pass:boolean
        }
    }
  },
  status: string
}>({} as {
  response: {
    Firstname: string,
    Lastname: string,
    Groups: Array<string>,
    Work: {
      prob1: {
        pass: boolean
      }
    }
  },
  status: string
})
const GetSelfData = async () => {
  try {
    const token = localStorage.getItem('token');
    if (!token?.length) throw new Error('Token empty or undefined');
    const response = await fetch(import.meta.env.VITE_api + "/api/getselfdata", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${token}`
      }
    });
    const json = await response.json();
    if (json.status === "success" && json.response) {
      SelfData.response = json.response;
      SelfData.status = json.status;
      return
    } else {
      SelfData.status = json.status;
    }
  } catch (error) {
    data.status = "error";
  }
}

onMounted(() => {
    GetAssigned();
    GetSelfData();
    GetGroupWork();
})
</script>
    
<template>
    <div class="Container">
        <div class="Assignmentwrap" v-if="data.response && data.response.length">
            <div class="PropositionHead">งานของฉัน</div>
            <div class="PropositionMenu" v-for="(Propo) in data.response">
                <div class="PropositionName" @click="$emit('SelectedChange', Propo )">{{ Propo.Name}}</div>
            </div>
        </div>
        <div class="Assignmentwrap" v-if="group_data.response && Object.keys(group_data.response).length">
            <div class="PropositionHead">งานของชั้นเรียน</div>
            <div class="PropositionMenu" v-for="(Propo, classname) in group_data.response" :key="classname">
                <div class="PropositionName"> ชั้นเรียน : {{ classname }}</div>
                <ul class="problemlist">
                    <li v-for="(problem, index) in Propo.Assigned" :key="index" 
                   
                    :class="`PropositionName ${ SelfData.response.Work[problem.Id]?.pass ? 'success' : '' }`" @click="$emit('SelectedChange', problem )">
                        {{ problem.Name}}</li>
                </ul>
            </div>
        </div>
    </div>
</template>

<style scoped>
.Container {
    z-index: 5;
    position: absolute;
    top: 10;
    left: 0;
    margin-top: 3rem;
    height: 82vh;
    overflow-x: hidden;
    transition: all 0.5s ease-out;
    background: var(--editor-PbarB);
    width: 0;
}
.Container.isshow{
    width: auto;
    min-width: 10rem;
    max-width: 15rem;
}
.PropositionMenu {
    padding: 0.3rem 1.5rem;
    position: relative;
    text-align: left;
    word-wrap: break-word;
}

.PropositionName {
    color: var(--editor-PbarF);
    font-size: 1rem;
    cursor: default;
    /*font-weight:1;
    margin-left:-1rem ;*/
}
.PropositionName.success{
    color: var(--editor-PbarSuccessF);
}
.PropositionHead{
    color: var(--editor-PbarF);
    font-size: 1.2rem;
    text-align: center;
    cursor: default;
    margin: 1rem 0;
}
.problemlist{
    margin-left: 2rem;
}
.problemlist > li {
    cursor: pointer;
}
.problemlist > li:hover{
    color: var(--editor-PbarF-H)
}
</style>
