<script setup lang="ts" >
import { reactive, onMounted } from 'vue';
import EditIcon from "./EditIcon.vue"
import TerminateIcon from "./TerminateIcon.vue"
import AddIcon from "./AddIcon.vue"

defineEmits(["TestcaseShowClick"])
const props = defineProps<{data: any}>()

const state = reactive({isEditing: false})
const editing = reactive<{data:any}>({data:{}})

onMounted(() => {
    editing.data = {...props.data}
    editing.data.Testcase = [...props.data.Testcase]
})

const addTestCaseHandler = () => {
    editing.data.Testcase.push({
        name:"Testcase"+(editing.data.Testcase.length+1),
        stdin:false,
        action:"add"
    })
}

const removeTestCaseHandler = (index:number) => {
    editing.data.Testcase[index].action = "remove";
}

const switchStdinHandler = (e: any, index:number) => {
    editing.data.Testcase[index].stdin = e.target?.checked
    editing.data.Testcase[index].action = "update"
} 

const outFileChangeHandler  = async (e: any, index:number) => {
    editing.data.Testcase[index].outfile = await fileToBase64(e.target.files[0])
    editing.data.Testcase[index].action = "update"
}

const inFileChangeHandler = async (e: any, index:number) => {
    editing.data.Testcase[index].infile = await fileToBase64(e.target.files[0])
    editing.data.Testcase[index].action = "update"
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

        const response = await fetch(import.meta.env.VITE_api + "/api/sentry/updatetestcase", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(
                {
                    id:editing.data.Id,
                    testcase:editing.data.Testcase
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
            <div class="exit-button" @click="$emit('TestcaseShowClick',null)">
                <svg viewBox="0 0 455.111 455.111" class="ColseIcon">
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
                {{ data.Name }} Testcase
                <div class="actionwrap">
                    <AddIcon v-if="state.isEditing" @click="addTestCaseHandler"/>
                    <EditIcon v-if="!state.isEditing" @click="() => state.isEditing = true" />
                    <TerminateIcon v-else @click="() => state.isEditing = false"/>
                </div>
            </div>
            <div class="testcasedetailwrapper">
                <div :class="`testcaselist ${Case.action === 'remove' ? 'remove' : ''}`" v-for="(Case,index) in editing.data.Testcase" :key="Case.name" v-if="state.isEditing" >
                    <div class="removewrap">
                        <TerminateIcon @click="removeTestCaseHandler(index)"/>
                    </div>
                    <div class="casetext"><span>Filename: </span>{{ Case.name }}</div>
                    <ul class="filelist">
                        <div>Files :</div>
                        <li>
                            <label>Input: {{ !Case.stdin ? "Not included" : ""}}</label>
                            <input type="file"  @change="(e) => inFileChangeHandler(e, index)" v-if="Case.stdin"/>
                        </li>
                        <li>
                            <label>Output: </label>
                            <input type="file" @change="(e) => outFileChangeHandler(e, index)"/>
                        </li>
                    </ul>
                    <div class="casetext">
                        <label>Use input: </label>
                        <input type="checkbox" :checked="Case.stdin" @change="(e) => switchStdinHandler(e,index)"/>
                    </div>
                </div>
                <div class="testcaselist" v-for="(Case,index2) in data.Testcase" :key="index2" v-else>
                    <div class="casetext"><span>Casename: </span>{{ Case.name }}</div>
                    <ul class="filelist">
                        <div>Files :</div>
                        <li v-if="Case.stdin">
                            <label>Input: {{ Case.name }}.in</label>
                        </li>
                        <li>
                            <label>Output: {{ Case.name }}.sol</label>
                        </li>
                    </ul>
                    <div class="casetext">
                        <label>Use input: </label>
                        <input type="checkbox" :checked="Case.stdin" disabled/>
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

.ColseIcon {
    width: 2.5rem;
    position: absolute;

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

.exit-button {
    position: absolute;
    top: 0;
    right: 0;
    cursor: pointer;
}

.Form {
    height: fit-content;
    border-radius: 1rem;
    text-align: left;
}

.FormInput {
    border-radius: 1rem;
    outline-color: var(--ad-modalOutline);
    height: 1.5rem;
    width: 80%;
    margin: 0.2rem 0;
    padding-left: 0.7rem;
    border-style: solid;
    text-align: left;
}

.FormInput:active,
.FormInput {
    border-color: var(--ad-modalOutline);
}

.FormButton {
    border: none;
    border-radius: 1rem;
    color: var(--ad-modalFormButF);
    background-color: var(--ad-modalFormButB);
    padding: 0.3rem 1rem;
    font-size: 1.25rem;
    text-align: center;
    width: 50%;
    cursor: pointer;
}

.controller {
    padding: 1rem;
    border-radius: inherit;
}

/*.formWrap {}*/

.inputWrapper {
    position: relative;
    margin-bottom: 1.5rem;
    text-align: left;
}

.inputbox {
    all: unset;
    outline: none;
    border-top: none;
    border-left: none;
    border-right: none;
    border-bottom: 1.75px solid;
    width: 100%;
}

#inputbox_file {
    margin-top: 0.5rem;
    padding-bottom: 0.5rem;
}

.placeholder {
    position: absolute;
    pointer-events: none;
    transition: all .2s;
    left: 0;
}

.focus .placeholder {
    transform: translateY(-100%) scale(0.9);
    font-size: 14px;
    left: -1%;
}

.inputbox:focus {
    border-bottom: 2px solid var(--ad-InputFocusB);
}

.loginbtn,
.signupbtn {
    all: unset;
    width: 100%;
    padding: 0.5rem 0;
    border-radius: 0.25rem;
    color: #fff;
    text-align: center;
    cursor: pointer;
    transition: background .2s;
}
.testcaselist{
    margin: 1rem 0;
    border: 1px solid black;
    padding: 1rem;
    position: relative;
}
.testcaselist.remove{
    filter: blur(0.25rem);
    pointer-events: none;
    cursor: not-allowed;
}

.filelist{
    list-style: none;
    margin: 0.25rem 0;
    margin-left: 1rem;
}
.testcasedetailwrapper{
    max-height: 30rem;
    overflow-y: auto;
}
.updatebuttonwrap{
    display: flex;
    justify-content: center;
    margin-top: 2rem;
}
.updatebuttonwrap > button{
    padding: 0.5rem 2rem;
}
.actionwrap{
    display: flex;
    gap: 0.5rem;
}
.removewrap{
    position: absolute;
    top: 3%;
    right: 1%;
}
</style>