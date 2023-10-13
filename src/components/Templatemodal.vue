<script setup lang="ts">
import { reactive, onMounted } from 'vue';
import EditIcon from "./EditIcon.vue"
import TerminateIcon from "./TerminateIcon.vue"
defineEmits(["TemplateHideClick"])
const props = defineProps<{data: any}>()

const state = reactive({isEditing: false})

const editing = reactive<{data:any}>({data:{}})

onMounted(() => {
    editing.data = {...props.data}
    editing.data.file = ""
    console.log(editing.data.Template)
})

const switchTemplateHandler = (e:any) => {
    editing.data.Template = e.target?.checked
    if(editing.data.Template === false){
        editing.data.action = "remove"
    }
}

const FileChangeHandler = async (e: any) => {
    editing.data.file = await fileToBase64(e.target.files[0])
    editing.data.action = "add"
}

const fileToBase64 = (file: File) => {
  return new Promise((resolve, reject) => {
    const Filereader = new FileReader();
    //@ts-ignore
    Filereader.onload = () => resolve(Filereader?.result?.split(',')[1]);
    Filereader.onerror = (error) => reject(error);
    Filereader.readAsDataURL(file);
  });
};


const updateClickHandler = async () => {
    try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api + "/api/sentry/updatetemplate", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(
                {
                    id:editing.data.Id,
                    action:editing.data.action,
                    file:editing.data.file
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

</script>

<template>
        <div class="Container">
        <div class="contentwrapper">
            <div class="exit-button" @click="$emit('TemplateHideClick',null)">
                <svg viewBox="0 0 455.111 455.111" class="CloseIcon">
                    <g>
                        <circle cx="227.556" cy="227.556" r="227.556" style="" fill="#cbcbcb" data-original="#e24c4b"
                            class="" opacity="1"></circle>
                        <path
                            d="M455.111 227.556c0 125.156-102.4 227.556-227.556 227.556-72.533 0-136.533-32.711-177.778-85.333 38.4 31.289 88.178 49.778 142.222 49.778 125.156 0 227.556-102.4 227.556-227.556 0-54.044-18.489-103.822-49.778-142.222 52.623 41.243 85.334 105.243 85.334 177.777z"
                            style="" fill="#838383" data-original="#d1403f" class="" opacity="1"></path>
                        <path
                            d="M331.378 331.378c-8.533 8.533-22.756 8.533-31.289 0l-72.533-72.533-72.533 72.533c-8.533 8.533-22.756 8.533-31.289 0-8.533-8.533-8.533-22.756 0-31.289l72.533-72.533-72.533-72.533c-8.533-8.533-8.533-22.756 0-31.289 8.533-8.533 22.756-8.533 31.289 0l72.533 72.533 72.533-72.533c8.533-8.533 22.756-8.533 31.289 0 8.533 8.533 8.533 22.756 0 31.289l-72.533 72.533 72.533 72.533c8.533 8.533 8.533 22.755 0 31.289z"
                            style="" fill="#ffffff" data-original="#ffffff" class=""></path>
                    </g>
                </svg>
            </div>
            <div class="HeaderText">
                {{ editing.data.Name }} Template
                <div class="actionwrap">
                    <EditIcon v-if="!state.isEditing" @click="() => state.isEditing = true" />
                    <TerminateIcon v-else @click="() => state.isEditing = false"/>
                </div>
            </div>
            <div class="templatecode-wrapper">
                <label>Template status:</label>
                <input type="checkbox" :checked="editing.data.Template" @change="(e) => switchTemplateHandler(e)" :disabled="state.isEditing === false"/>
                <div class="templatecode-file-wrap">
                    <label>Template: {{ editing.data.Template ? editing.data.Id+".tp" : "Not included" }}</label>
                    <div class="inputfile-wrapper" v-if="state.isEditing">
                        <input type="file" @change="(e) => FileChangeHandler(e)"/>
                    </div>
                </div>
            </div>
            <div class="updatebuttonwrap" v-if="state.isEditing">
                <button @click="updateClickHandler">Update</button>
            </div>
        </div>
    </div>
</template>

<style scoped>
.Container {
    width: 100%;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 100;
    backdrop-filter: blur(0.5rem);
}

.CloseIcon {
    width: 2.5rem;
    position: absolute;

}
.exit-button {
    position: absolute;
    top: 0;
    right: 0;
    cursor: pointer;
}

.HeaderText {
    font-size: 2rem;
    font-weight: 700;
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 2rem;
}
.contentwrapper {
    padding: 2rem;
    background-color: var(--ad-modalIconWrapB);
    border-radius: 2rem;
    font-size: 1.2rem;
    position: relative;
    max-width: 30rem;
}

.templatecode-file-wrap{
    margin-top: 1rem;
}
.actionwrap{
    display: flex;
    gap: 0.5rem;
}
.updatebuttonwrap{
    display: flex;
    justify-content: center;
    margin-top: 2rem;
}
.updatebuttonwrap > button{
    padding: 0.5rem 2rem;
}
</style>