<script setup lang="ts">
import LoadingIcon from "../components/LoadingIcon.vue"
import FooterBar from "../components/FooterBar.vue"
import NavigationBar from '../components/NavigationBar.vue'
import PropositBar from '../components/PropositBar.vue'
import ComplieRequest from "../module/requestcompile"
// import PassPopup from "../components/PassPopup.vue"
import DisplayIcon from "../components/DisplayIcon.vue"
import NonepassPopup from "../components/NonepassPopup.vue"
import MenuIcon from "../components/MenuIcon.vue"
import { Codemirror } from 'vue-codemirror'
import { cpp, cppLanguage } from '@codemirror/lang-cpp'
import { oneDark, light } from '../module/editorTheme'
import {snippetCompletion} from "@codemirror/autocomplete"
import { ref, reactive, computed, onMounted } from 'vue';
import useUserStore from "../module/userstore"
import useTheme from "../module/usetheme"
import { useRoute } from "vue-router"

const userdata = useUserStore();
const themedata = useTheme();
const inputRef = ref();
const FontSizeValue = reactive<{Value:number}>({Value:100})

const state = reactive<{ awaitsubmit: boolean, showresultside: boolean, showcodeside: boolean }>({
  awaitsubmit: false,
  showcodeside: true,
  showresultside: true
})

type OutputStructure = { case: string, expected: string, output: string, status: string }

const NonepasspopupState = reactive<{ IsShow: boolean, output: Array<OutputStructure> }>({
  IsShow: false,
  output: []
})
type ArrLength = Array<{ [arrayLenght: number]: { name: "", stdin: false } }>
type ProblemData = { Name: string, Id: string, Link: string, Info: string, Testcase: ArrLength, Template:boolean }
// type SelectedProblem = { Problem: ProblemData }
const UserCode = reactive<{Name:any,IsShow:boolean}>({Name:'',IsShow:false})

const ProbSelected = reactive<{ Problem: ProblemData }>({ Problem: { } as ProblemData })
const MenuTabIsShow = reactive<{ IsShow: boolean }>({
  IsShow: false
})

const data = reactive<{ code: string, output: string, input: string, status: string }>({
  code: "",
  input: "",
  output: "",
  status: ""
});
const view = ref()

const extensions = reactive<{data:Array<any>}>({data : [
  cpp(),
  cppLanguage.data.of({
  autocomplete: [
    snippetCompletion('int ${varname}', {label: 'int', detail: 'Integer'}),
    snippetCompletion('long ${next} ${type} ${varname}', {label: 'long', detail: 'Increasing bit'}),
    snippetCompletion('float ${varname}', {label: 'float', detail: 'Float'}),
    snippetCompletion('int ${varname}', {label: 'double', detail: 'Double'}),
    snippetCompletion('string ${varname}', {label: 'string', detail: 'String'}),
    snippetCompletion('#include<${headerfile}>', {label: '#include', detail:'Add header file'}),
    snippetCompletion('#define ${name} ${value};', {label: '#define', detail: 'Define variable'}),
    snippetCompletion('using namespace ${class};', {label: 'using namespace', detail: 'Shorter namespace'}),
    snippetCompletion('int main(int argc, char** argv){\n\n\treturn 0;\n}', {label: 'int main', detail: 'Main function'}),
    snippetCompletion('while(${condition}){\n\n}', {label: 'while', detail: 'While loop'}),
    snippetCompletion('do{\n\n}while(${condition});', {label: 'do while', detail: 'Do While loop'}),
    snippetCompletion('for(${initial}; ${condition}; ${executor}){\n\n}', {label: 'for', detail: 'For loop'}),
    snippetCompletion('${returntype} ${functionname}(${param}){\n\n}', {label: 'function', detail: 'Create function'}),
    snippetCompletion('class ${classname}{\n\n};', {label: 'class', detail: 'Create class'}),
    snippetCompletion('if(${condition}){\n\n}', {label: 'if', detail: 'If control'}),
    snippetCompletion('else if(${condition}){\n\n}', {label: 'else if', detail: 'Else if control'}),
    snippetCompletion('else{\n\n}', {label: 'else', detail: 'Else control'}),
    snippetCompletion('switch(${lookup}){\n\tcase ${matchcase}:\n\n\tdefault:\n\n}', {label: 'switch', detail: 'switch control'}),
    snippetCompletion('#include<iostream>\nusing namespace std;\n\nint main(int argc, char** argv){\n\n\treturn 0;\n}', {label: 'Init', detail: 'Initial code'}),
  ]
})

]})

const getcurrentExtension = computed(() => themedata.theme === 'dark' ? [...extensions.data, oneDark] : [...extensions.data, light])

const lines = computed(() => typeof data.output === 'string' ? data.output.split(/\r?\n/) : []);

const onSelectedChange = (Problem: ProblemData) => {
  MenuTabClick();
  ProbSelected.Problem = Problem;
  if(ProbSelected.Problem.Template){
    GetTemplateCode();
  }
}

const GetTemplateCode = async () => {
  try {
      const token = localStorage.getItem('token');
      if (!token?.length) throw new Error('Token empty or undefined');

      const response = await fetch(import.meta.env.VITE_api + "/api/gettemplatecode?sol=" + ProbSelected.Problem.Id, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`
        }
      });
      const json = await response.json();
      if (json.status === "success" && json.response) {

        data.code = json.response;
        data.status = json.status;
        console.log(data)
        return
      } else {
        data.status = json.status;
      }
  } catch (error) {
    data.status = "error";
  }
}

const handleReady = (payload: {
  view: import("@codemirror/view").EditorView;
  state: import("@codemirror/state").EditorState;
  container: HTMLDivElement;
}) => {
  view.value = payload.view
}




const onRunClick = async () => {
  if (state.awaitsubmit) return;
  data.output = "";
  if (!state.showresultside) {
    state.showresultside = true
  }
  state.awaitsubmit = true;
  const response = await ComplieRequest(data.code, data.input, "c++", false);
  state.awaitsubmit = false;
  if (response?.status && response?.output) {
    data.output = response.output;
    data.status = response.status;
  }
}

const onGraderSubmit = async () => {
  if (state.awaitsubmit) return;
  if (!state.showresultside) {
    state.showresultside = true
  }

  data.output = "";
  state.awaitsubmit = true;
  const response = await ComplieRequest(data.code, data.input, "c++", true, ProbSelected.Problem.Id, ProbSelected.Problem.Testcase);
  state.awaitsubmit = false;
  if (response?.status && response?.output) {
    data.output = response.output;
    data.status = response.status;
  }

  if (response && typeof response?.output === "object") {
    NonepasspopupState.output = response.output;
    state.awaitsubmit = false;
    NonepasspopupState.IsShow = true;
  }
}

const CloseClick = () => {
  NonepasspopupState.IsShow = false;
}

const onFileChange = (event: any) => {
  const file = event.target?.files[0];
  if (file) ReadFile(file);
}

const ReadFile = (file: any) => {
  const fileReader = new FileReader();
  fileReader.onload = () => {
    data.code = typeof fileReader.result === "string" ? fileReader.result : "";
  }
  fileReader.readAsText(file);
}

const MenuTabClick = () => {
  MenuTabIsShow.IsShow = !MenuTabIsShow.IsShow;
}

const onSubmitClick = () => {
  if (inputRef.value && inputRef.value.click)
    inputRef.value.click();
}



const route = useRoute();
onMounted(async () => {
  let ProbsParam = route.params.sol;
  let UserParam = route.params.target;
  let UserFNameParam = route.params.fname;
  let UserLNameParam = route.params.lname;
  if (userdata.info?.isadmin && ProbsParam !== undefined && UserParam !== undefined && UserFNameParam!== undefined && UserLNameParam!== undefined) {
    UserCode.Name = UserFNameParam+" "+UserLNameParam;
    try {
      const token = localStorage.getItem('token');
      if (!token?.length) throw new Error('Token empty or undefined');

      const response = await fetch(import.meta.env.VITE_api + "/api/sentry/getusercode/?sol=" + ProbsParam + "&target=" + UserParam, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`
        }
      });
      const json = await response.json();
      if (json.status === "success" && json.response) {

        data.code = json.response;
        data.status = json.status;
        console.log(data)
        return
      } else {
        data.status = json.status;
      }
    } catch (error) {
      data.status = "error";
    }
  }
}
)


</script>

<template>
  <NonepassPopup :output="NonepasspopupState.output" v-if="NonepasspopupState.IsShow" @CloseClick="CloseClick" />
  <NavigationBar />
  <PropositBar :class="`${MenuTabIsShow.IsShow ? 'isshow' : ''}`" @SelectedChange="onSelectedChange" />
  <div class="Menu">
    <div style="  margin-left: 1rem;">
      <MenuIcon @click="MenuTabClick" v-if="userdata.info" />
    </div>
    <div class="PropoHead" style="font-size: 1.2rem;">{{ ProbSelected.Problem.Name }}
      <a :href="ProbSelected.Problem.Link" target="_blank" rel="noopener noreferrer"
        v-if="ProbSelected.Problem.Link && ProbSelected.Problem.Link.length">
        (ดูโจทย์)
        <!-- <div class="controller-buttonwrapper"> -->
      </a>
      <a :href="ProbSelected.Problem.Info" target="_blank" rel="noopener noreferrer"
        v-if="ProbSelected.Problem.Info && ProbSelected.Problem.Info.length">
        (ดูโค๊ด)
        <!-- <div class="controller-buttonwrapper"> -->
      </a>
      <div v-if="UserCode.Name && UserCode.Name.length" style="font-size: 1.2rem;margin-left: 1rem;">
       Code from : {{ UserCode.Name }}
      </div>
    </div>
    
    <button @click="onSubmitClick()" :class="`controller-button `" v-if="userdata.info">
      <div class="inputfileButton">
        <svg viewBox="0 0 494.554 494.554">
          <g>
            <path
              d="M464.962 16.626C464.962 7.441 457.507 0 448.337 0H114.845c-9.171 0-16.626 7.441-16.626 16.626v141.807h33.252V33.252H431.71v351.659h-59.795c-9.176 0-16.625 7.436-16.625 16.627v59.766H131.471V337.159H98.219v140.769c0 9.186 7.455 16.626 16.626 16.626h257.07c4.593 0 8.749-1.87 11.752-4.873l.016-.016 76.391-76.359.016-.016a16.56 16.56 0 0 0 4.873-11.752V16.626z"
              fill="#fff" data-original="#fff" class=""></path>
            <path
              d="M172.989 327.305a14.522 14.522 0 0 0 8.977 13.428 14.565 14.565 0 0 0 5.569 1.104c3.798 0 7.519-1.474 10.292-4.255l79.492-79.508a14.549 14.549 0 0 0 4.272-10.277c0-3.865-1.544-7.563-4.272-10.291l-79.492-79.492a14.524 14.524 0 0 0-10.292-4.27c-1.881 0-3.767.359-5.569 1.101a14.564 14.564 0 0 0-8.977 13.445v23.396H44.139c-8.022 0-14.548 6.508-14.548 14.547v83.129c0 8.037 6.525 14.549 14.548 14.549h128.85v23.394z"
              fill="#fff" data-original="#fff" class=""></path>
          </g>
        </svg>
        <label>Import</label>
      </div>
    </button>

    <button @click="onRunClick()" :class="`controller-button ${state.awaitsubmit ? 'submiting' : ''}`">Run</button>
    <input ref="inputRef" type="file" class="FileInput" id="FileInput" style="display:none;" accept=".cpp"
      @change="onFileChange">

    <button @click="onGraderSubmit" :class="`controller-submit-button  ${state.awaitsubmit ? 'submiting' : ''}`"
      v-if="userdata.info && ProbSelected.Problem.Id && ProbSelected.Problem.Id.length">
      Submit Code
    </button>
    <button v-else @click="MenuTabClick" class="controller-submit-button">
      เลือกโจทย์
    </button>
  </div>
  <div class="size-controller-wrapper">
    <div class="MiddleBar"></div>
    <div class="size-controller">
      <div class="controller">
        <DisplayIcon :display="state.showcodeside" @click="() => state.showcodeside = !state.showcodeside" />
        <div class="controller-text">Code</div>
        <div class="input-ramge-controller">
          <svg viewBox="0 0 522.781 522.782" class="ZoomIcon">
              <g>
                <path
                  d="m512.703 455.395-75.353-75.2c-7.191-7.19-16.983-10.557-26.469-10.098l-36.262-36.261c61.735-81.625 55.463-198.365-18.896-272.799-81.318-81.319-213.28-81.319-294.599.153-81.473 81.472-81.473 213.282-.153 294.601 74.435 74.435 191.174 80.631 272.799 18.896l36.261 36.262c-.459 9.485 2.907 19.277 10.098 26.469l75.2 75.199c13.464 13.464 35.419 13.464 48.883 0l8.339-8.339c13.616-13.464 13.616-35.266.152-48.883zM310.666 310.733c-56.457 56.457-148.18 56.61-204.637.153s-56.457-148.027.153-204.638c56.457-56.457 148.181-56.61 204.638-.153s56.227 148.028-.154 204.638z"
                  fill="#fff" data-original="#fff" class=""></path>
                <path
                  d="M265.072 186.727h-34.119v-34.119c0-12.316-9.945-22.262-22.108-22.262h-.229c-12.164 0-22.108 9.945-22.108 22.262v34.119h-34.119c-12.316 0-22.261 9.945-22.261 22.109v.153c0 12.164 9.945 22.108 22.261 22.108h34.119v34.119c0 12.317 9.945 22.262 22.108 22.262h.153c12.163 0 22.108-9.944 22.108-22.262v-34.119h34.119c12.316 0 22.262-9.945 22.262-22.108v-.153c.152-12.164-9.792-22.109-22.186-22.109z"
                  fill="#fff" data-original="#fff" class=""></path>
              </g>
            </svg>
          <div class="slidecontainer">
            <input type="range" min="80" max="200" v-model="FontSizeValue.Value" class="slider" id="myRange"/>
          </div>
        </div>
      </div>
      <div class="controller">
        <div class="controller-text">Result</div>
        <DisplayIcon :display="state.showresultside" @click="() => state.showresultside = !state.showresultside" />
      </div>
    </div>
    <div class="MiddleBar"></div>
  </div>
  <div class="content">
    <div class="MiddleBar"></div>
    <codemirror v-if="state.showcodeside" v-model="data.code" placeholder="C++ Code here..." class="codewrapper"
      v-bind:style="{ height: '80vh', width: '100%', gridColumn: state.showresultside ? '' : 'span 3', maxWidth: state.showresultside ? '50vw' : '100vw', fontSize: FontSizeValue.Value*0.015+'rem'  }"
      :autofocus="true" :indent-with-tab="true" :tab-size="4" :extensions="getcurrentExtension" @ready="handleReady" />
    <div class="MiddleBar" v-if="state.showcodeside"></div>
    <div class="result-container" v-if="state.showresultside" :style="state.showcodeside ? '' : 'grid-column: span 3;'">
      <div class="stdin-inputwrap" :style="data.input.split('\n').length>4? 'height: 50%' : '20%'">
        <h4>Input : </h4>
        <textarea class="stdin" v-model="data.input" ></textarea>
      </div>
      <div class="output-wrap">
        <div class="loadingwrap" v-if="state.awaitsubmit">
          <LoadingIcon />
        </div>
        <div v-for="line in lines" :key="line" :class="`output ${data.status !== 'success' ? 'error' : ''}`">{{ line }}
        </div>
      </div>
    </div>
    <div class="MiddleBar"></div>
  </div>
  <FooterBar />
</template>

<style scoped>
.content {
  display: grid;
  grid-template-columns: 1rem 1fr 1rem 1fr 1rem;
}

.Menu {
  width: 100%;
  height: 3rem;
  background-color: var(--edit-menuB);
  display: flex;
  align-items: center;
  justify-content: space-between;
  /*padding: 1rem;*/
}

.PropoHead {
  color: var(--editor-F);
  margin-left: 1rem;
  width: 38vw;
  display: flex;
  gap: 1rem;
}

.PropoHead>a {
  color: var(--editor-F);
  text-decoration: underline;
}

.PropoHead>a ::-webkit-textfield-decoration-container {
  color: var(--editor-F)
}

.controller-container {
  padding: 1rem;
  background: var(--editor-conB);
  display: flex;
  flex-direction: column;
  align-items: center;

}

.size-controller-wrapper {
  display: grid;
  grid-template-columns: 1rem 1fr 1rem;
}
.ZoomIcon{
  margin-left: 1rem;
  margin-right: 0.5rem;
  width: 1.3rem;
  aspect-ratio: 1/1;
  margin-top: 0.2rem;
}
.ZoomIcon path{
  fill: var(--editor-zoomIcon-C);
}
.slidecontainer{
  display:flex;
}
.slider {
  -webkit-appearance: none;
  width: 100%;
  background: var(--editor-zoomB);
  outline: none;
  opacity: 0.7;
  margin-top: 0.3rem;
  -webkit-transition: .2s;
  transition: opacity .2s;
  border-radius: 1rem;
  overflow: hidden;
}

.slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 1rem;
  aspect-ratio: 1/1;
  background: var(--editor-zoomI-B);
  cursor: pointer;
}

.slider::-moz-range-thumb {
  background: var(--editor-zoomI-B);
  cursor: pointer;
}

.size-controller {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  padding: 0 1rem;
}

.controller {
  display: flex;
  gap: 0.5rem;
  color: var(--editor-conF);
}

.controller-text {
  display: flex;
  align-items: center;
  font-weight: bold;
}
.input-ramge-controller{
  display: flex;
  align-items: center;
  font-weight: bold;
  max-width: 10vw;
}
.controller-submit-button {
  margin: 0 1rem 0 auto;
  padding: 0.5rem 2rem;
  border: none;
  background: var( --editor-conMainB);
  color: var(--editor-conMainF);
  border-radius: 1rem;
  box-shadow: 0 0 0.25rem 0.1rem var(--editor_conShadow);
  transition: all .25s ease-in-out;
  cursor: pointer;
  font-weight: bold;
  word-wrap:nowrap;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.controller-submit-button-else {
  margin: 0 1rem 0 auto;
}

.controller-button {
  padding: 0.5rem 2rem;
  border: none;
  background: var( --editor-conMainB);
  color: var(--editor-conMainF);
  border-radius: 1rem;
  box-shadow: 0 0 0.25rem 0.1rem var(--editor_conShadow);
  transition: all .25s ease-in-out;
  cursor: pointer;
  font-weight: bold;
  margin: 0 0.4rem 0 1rem;
}

.controller-button.submiting, .controller-submit-button.submiting {
  cursor: not-allowed;
}

.controller-button:hover, .controller-submit-button:hover {
  background: var( --editor-conMainB-H);
  color: var(--editor-conMainF-H);
  box-shadow: 0 0 0.25rem 0.1rem var(--editor_conShadow-H);
}

.controller-button label {
  font-size: 1rem;
  margin-left: 0.2rem;
}

.controller-button:hover label {
  color: var(--editor-conF-H);
  cursor: pointer;
}

.controller-button:hover path {
  fill: var(--editor-conB-H);
}

.MiddleBar {
  background-color: var(--editor-midBarB);
}

.result-container {
  padding: 1rem;
  color: var(--editor-ResF);
  font-weight: bold;
  background: var(--editor-ResB);
  height: 80vh;
  overflow-y: auto;
}

.output-wrap {
  position: relative;
  width: 100%;
  height: 80%;
}

.output {
  margin-bottom: 1rem;
  color: var(--editor-outF);
}

.output.error {
  color: var(--editor-errResF);
}

.stdin-inputwrap {
  min-height: 20%;
  max-height: 50%;
  display: flex;
  flex-direction: column;
}

.stdin {
  width: 100%;
  height: 100%;
  resize: none;
  color: var(--editor-stdinF);
  font-size: 1.25rem;
  background: var(--editor-stdinB);
  flex-grow: 1;
  margin-bottom: 1rem;
}

.loadingwrap {
  width: 100%;
  height: 100%;
  position: absolute;
  display: flex;
  justify-content: center;
  align-items: center;
}

.inputfileButton {
  text-align: center;
  color: var(--editor-inputF);
  align-items: center;
  text-align: center;
  display: flex;
}

.inputfileButton:hover {
  color: var(--editor-inputF-H);
}

.inputfileButton>svg {
  /* width: 100%; */
  height: 1.1rem;
  aspect-ratio: 1/1;
  margin-right: 0.1rem;
  cursor: pointer;
}
</style>
