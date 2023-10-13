<script setup lang="ts">
// import LoadingIcon from "../components/LoadingIcon.vue"
// import FooterBar from "../components/FooterBar.vue"
import { useUserStore } from "../module/userstore"
import NavigationBar from '../components/NavigationBar.vue'
import GroupTable from '../components/GroupTable.vue'
import { reactive, onMounted, ref, computed } from 'vue';
import { useRouter } from 'vue-router'
import UserTable from '../components/UserTable.vue'
import UserTableByGroup from '../components/UserTableByGroup.vue'
import AddGroupPopup from "../components/AddGroupPopup.vue"
import AddAssignment from "../components/AddAssignment.vue"
import AssignmentTable from "../components/AssignmentTable.vue"
import EditProbPopup from "../components/EditProbPopup.vue";
import Testcasemodal from "../components/Testcasemodal.vue"
import UserGroupTable from "../components/UserGroupTable.vue";
import AddUsermodal from "../components/AddUsermodal.vue";
import UpdateUsermodal from "../components/UpdateUsermodal.vue";
import Templatemodal from "../components/Templatemodal.vue";
import { ExportCSV } from "../module/exporter";
const userdata = useUserStore();
const router = useRouter();
const DropDownRef = ref()
const ButtonRef = ref()
const Page = reactive<Page>({
    selected: 0,
})
const Dropdown = reactive<{ Isshow: boolean }>({
    Isshow: false,
})
const Assingment = reactive<{ Isshow: boolean }>({
    Isshow: false,
})
const addUser = reactive<{ Isshow: boolean }>({
    Isshow: false,
})
const GroupData = reactive<{response:{[key:string]:any}, status:string}>({} as {response:{[key:string]:any}, status:string})

const Filter = reactive<{state:string | number}>({state:"ALL"});

const computedGroupData = computed(() => GroupData.response[Filter.state])

const DropdownShow = () => {
    Dropdown.Isshow = true;
    window.addEventListener("click", outSideDetected)
}
const outSideDetected = (e: Event) => {
    if (!DropDownRef.value) {
        Dropdown.Isshow = false;
        window.removeEventListener("click", outSideDetected)
    }
    if (!DropDownRef.value.contains(e.target) && !ButtonRef.value.contains(e.target)) {
        Dropdown.Isshow = false;
        window.removeEventListener("click", outSideDetected)
    }
}
const AddGroupState = reactive<{ Isshow: boolean }>({
    Isshow: false,
})


const EditProbPopupState = reactive<{ Isshow: boolean, Prob: object }>({
    Isshow: false,
    Prob: {}
})

const onMenuButtonClick = (pagenum: number) => {
    Page.selected = pagenum;
}

const onAddGroupBtnClick = () => {
    AddGroupState.Isshow = !AddGroupState.Isshow;
}

const onAddAssingBtnClick = () => {
    Assingment.Isshow = !Assingment.Isshow;
}
const onAddUserBtnClick = () => {
    addUser.Isshow = !addUser.Isshow;
}

const onEditProbPopupBtnClick = (Prob: object) => {
    EditProbPopupState.Isshow = !EditProbPopupState.Isshow;
    EditProbPopupState.Prob = Prob
}


const EditUsermodalState = reactive<{
    IsShow: Boolean,
    Username:string
}>({
    IsShow: false,
    Username:''
})

const onUpdateUserClick = (Username:string) => {
    EditUsermodalState.Username = Username;
    EditUsermodalState.IsShow = !EditUsermodalState.IsShow;
}

const onEditUserClick = () => {
    EditUsermodalState.IsShow = !EditUsermodalState.IsShow;
}



type TestcaseType = {
    Id: string,
    Name: string,
    Info: string,
    Link: string,
    Testcase: any
    Template: boolean
}

const Testcase = reactive<{ Isshow: boolean, Data: TestcaseType }>({} as { Isshow: boolean, Data: TestcaseType })

const TestcaseShowClick = (TC: any) => {
    Testcase.Isshow = !Testcase.Isshow
    console.log(Testcase.Data)
    Testcase.Data = TC;

}

type TemplateType = TestcaseType

const Template = reactive<{ Isshow: boolean, Data: TemplateType }>({} as { Isshow: boolean, Data: TemplateType })

const TemplateShowClick = (TC: any) => {
    Template.Isshow = !Template.Isshow
    console.log(Template.Data)
    Template.Data = TC;

}

const TemplateHideClick = () => {
    Template.Isshow = false;
    Template.Data = {} as TemplateType;
}

const GetGroupData = async () => {
    try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api + "/api/sentry/getgroupdata", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            }
        });

        const json = await response.json();
        if (json.status === "success" && json.response) {
            GroupData.response = json.response;
            GroupData.status = json.status;
        }
    } catch (error) {
        GroupData.status = "error";
    }
}

const UserTableData = reactive<TableData>(
    {} as TableData
)


const GetUsertable = async () => {
    try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api + "/api/sentry/getusertable", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            }
        });
        const json = await response.json();
        if (json.status === "success" && json.response) {
            UserTableData.response = json.response;
            UserTableData.status = json.status;
        } else {
            UserTableData.status = json.status;

        }
    } catch (error) {
        UserTableData.status = "error";
    }
}

const PrepareDataToExport = () => {
    const _U = {
        "NAME": "ชื่อ-สกุล",
        "PASS": "ผ่าน",
        "NOTPASS": "ไม่ผ่าน",
        "NOTSUBMIT": "ไม่ได้ส่ง",
        "MAX":"ทั้งหมด",
        "CURRENT":"ทำแล้ว"
    }
    if(Filter.state === 'ALL'){
        const processeddata:any = {} 
        const Assignments_Id = Object.keys(UserTableData.response.Structure)
        const Assignments_Count = Assignments_Id.length
        processeddata.header = [_U['NAME'], _U['CURRENT'], _U['MAX'], ...Assignments_Id.map(id => UserTableData.response.Structure[id].Name + `(${UserTableData.response.Structure[id].Maxscore})`)]
        processeddata.row =  UserTableData.response.Assigned.map((item) => {
            const rowdata:any = {
                name: item.Firstname + " " + item.Lastname,
                current: item.Score,
                max:Assignments_Count
            }
            for(const key of Assignments_Id){
                if(item.Data[key]){
                    rowdata[key] = item.Data[key].Pass ? _U['PASS']+ " | attempt "+ item.Data[key].Attempt + " | at "+ (new Date(item.Data[key].Timestamp * 1000).toLocaleString()).split(',').join(' ') : _U['NOTPASS'] + " | " + item.Data[key].Score+ " | attempt "+ item.Data[key].Pass + " | at "+ (new Date(item.Data[key].Timestamp * 1000).toLocaleString()).split(',').join(' ')
                }else{
                    rowdata[key] = _U['NOTSUBMIT']
                }
            }
            return rowdata;
        })
        ExportCSV(processeddata)
    }else{
        const processeddata:any = {} 
        const SelectedGroupData = GroupData.response[Filter.state]
        const Assignments_Id = Object.keys(SelectedGroupData.Assignment)
        const Assignments_Count = Assignments_Id.length
        processeddata.header = ["ชื่อ-สกุล", _U['CURRENT'], _U['MAX'], ...Assignments_Id.map(id => SelectedGroupData.Assignment[id].Name + `(${SelectedGroupData.Assignment[id].Maxscore})`)]
        processeddata.row = Object.keys(SelectedGroupData.Member).map((user) => {
            const item = SelectedGroupData.Member[user]
            let count = 0;
            Assignments_Id.forEach((data: string) => {
                const p = SelectedGroupData.Member[user].Work;
                if (p[data] && p[data].Pass) {
                    count++;
                }
            })
            const rowdata:any = {
                name: item.Firstname + " " + item.Lastname,
                score: count,
                max:Assignments_Count
            }
            for(const key of Assignments_Id){
                if(item.Work[key]){
                    rowdata[key] = item.Work[key].Pass ? _U['PASS']+" | attempt "+  item.Work[key].Attempt + " | at "+ (new Date(item.Work[key].Timestamp * 1000).toLocaleString()).split(',').join(' ') : _U['NOTPASS'] + " | " + item.Work[key].Score + " | attempt "+  item.Work[key].Attempt +" | at "+ (new Date(item.Work[key].Timestamp* 1000).toLocaleString()).split(',').join(' ')
                }else{
                    rowdata[key] = _U['NOTSUBMIT']
                }
            }
            return rowdata;
        })
        ExportCSV(processeddata)
     
    }
}

onMounted(async () => {
    if (!userdata.info?.isadmin) {
        router.push("/editor");
        return
    }
    GetUsertable()
    GetGroupData();
})
</script>

<template >
    <!-- <div style="width:100vw;height:100vh;overflow: hidden"> -->
    <AddGroupPopup v-if="AddGroupState.Isshow" @onAddGroupBtnClick="onAddGroupBtnClick" />
    <AddAssignment v-if="Assingment.Isshow" @onAddAssingBtnClick="onAddAssingBtnClick" />
    <Testcasemodal v-if="Testcase.Isshow" :data="Testcase.Data" @TestcaseShowClick="TestcaseShowClick" />
    <Templatemodal v-if="Template.Isshow" :data="Template.Data" @TemplateHideClick="TemplateHideClick" />
    <AddUsermodal v-if="addUser.Isshow" @onAddUserBtnClick="onAddUserBtnClick" />
    <UpdateUsermodal v-if="EditUsermodalState.IsShow" :userName="EditUsermodalState" @onEditUserClick="onEditUserClick"/>
    <EditProbPopup v-if="EditProbPopupState.Isshow" :Probs="EditProbPopupState.Prob"
        @onEditProbPopupBtnClick="onEditProbPopupBtnClick"  />
    <NavigationBar />
    <div class="MenuContainer">
        <button type="button" class="MenuButton" @click="onMenuButtonClick(0)"
            :style="` ${Page.selected === 0 ? 'color: white;cursor:text;' : 'color: gray;cursor:pointer;'}`">User</button>
        <button type="button" class="MenuButton" @click="onMenuButtonClick(1)"
            :style="` ${Page.selected === 1 ? 'color: white;cursor:text;' : 'color: gray;cursor:pointer;'}`">User
            Group</button>
        <button type="button" class="MenuButton" @click="onMenuButtonClick(2)"
            :style="` ${Page.selected === 2 ? 'color: white;cursor:text;' : 'color: gray;cursor:pointer;'}`">Assignment</button>
        <button type="button" class="MenuButton" @click="onMenuButtonClick(3)"
            :style="` ${Page.selected === 3 ? 'color: white;cursor:text;' : 'color: gray;cursor:pointer;'}`">Group</button>
    </div>
    <div class="exportwrapper" v-if="Page.selected === 0">
        <button class="controller-button" @click="() => PrepareDataToExport()">Export CSV</button>
    </div>
    <div class="MenuContainer" v-if="Page.selected === 0">
        <button type="button" class="MenuButton" @click="Filter.state = 'ALL'" :style="`${Filter.state === 'ALL' ? 'color: white;cursor:text;' : 'color: gray;cursor:pointer;'}`">All</button>
        <button v-for="(_, index) in GroupData.response" type="button" class="MenuButton" :style="`${Filter.state === index ? 'color: white;cursor:text;' : 'color: gray;cursor:pointer;'}`" @click="Filter.state = index">{{index}}</button>
    </div>
    <div style="" class="MainContainer" :style="`--left:-${Page.selected * 100}vw;`">
        <div class="wraptable">
            <div class="container">
                <div class="TableContainer">
                    <UserTable @onUpdateUserClick="onUpdateUserClick" :tabledata="UserTableData" :key="UserTableData.status" v-if="Filter.state === 'ALL'" />
                    <UserTableByGroup :tabledata="computedGroupData" :key="Filter.state" v-else/>
                </div>
            </div>
            <div class="container">
                <div class="TableContainer">
                    <UserGroupTable />
                </div>
            </div>
            <div class="container">
                <div class="TableContainer">
                    <AssignmentTable 
                        @onEditProbPopupBtnClick="onEditProbPopupBtnClick"
                        @TestcaseShowClick="TestcaseShowClick" 
                        @TemplateShowClick="TemplateShowClick" />
                </div>
            </div>
            <div class="container">
                <div class="TableContainer">
                    <GroupTable />
                </div>
            </div>

        </div>
        <div>

        </div>
        <div ref="ButtonRef" class="AddGroupBtn" @click="DropdownShow">
            <svg viewBox="0 0 469.333 469.333">
                <g>
                    <path
                        d="M437.332 192h-160V32c0-17.664-14.336-32-32-32H224c-17.664 0-32 14.336-32 32v160H32c-17.664 0-32 14.336-32 32v21.332c0 17.664 14.336 32 32 32h160v160c0 17.664 14.336 32 32 32h21.332c17.664 0 32-14.336 32-32v-160h160c17.664 0 32-14.336 32-32V224c0-17.664-14.336-32-32-32zm0 0"
                        fill="#ff8800" data-original="#ff8800" class=""></path>
                </g>
            </svg>
            <ul ref="DropDownRef" class="DropdownContainer" v-if="Dropdown.Isshow">
                <li class="DropdownMenu" @click="onAddGroupBtnClick">Add group</li>
                <li class="DropdownMenu" @click="onAddUserBtnClick">Add User</li>
                <li class="DropdownMenu" @click="onAddAssingBtnClick">Add Assignment</li>
            </ul>
        </div>
    </div>
    <!-- </div> -->
</template>
<style scoped>
.MenuContainer {
    background-color: var(--ad-menuconB);
}

.Icon-wrap {
    max-height: 2rem;
}

.MenuButton {
    padding: 0.5rem 1rem;
    border-radius: 0.5rem 0.5rem 0 0;
    border: none;
    margin: 0 0.2rem;
    color: gray;
    background-color: var(--ad-menuButB)
}

.MainContainer {
    width: 100vw;
    overflow: hidden;
}

.wraptable {
    width: 100%;
    display: flex;
    transform: translate(var(--left));
    transition: all 0.5s ease-out;
    ;
}

.container {
    width: 100%;
    height: 85vh;
}

.TableContainer {
    width: 100%;
    height: 85vh;
    overflow: scroll;
    margin: 0 auto;
}

.Table {
    border-collapse: collapse;
    /*border: 2px solid #ddd;*/
    font-size: 1.2rem;
    /*table-layout: auto;*/
    /* ตั้งค่าให้ตารางมีขนาดคงที่ */
    overflow-y: scroll;
    overflow-x: scroll;
    text-align: center;
}

.TableContainer {
    width: 100vw;
}

.TableContainer::-webkit-scrollbar {
    width: 0.35rem;
    height: 0.15rem;
}

.TableContainer::-webkit-scrollbar-thumb {
    background: #4e4e4e;
}

/*.TableContainer::-webkit-scrollbar-track{}*/
.TableContainer::-webkit-scrollbar-corner {
    background: transparent;
    width: 0;
    height: 0;
}

.Table:deep() th,
.Table:deep() td {
    /*border: 2px solid #ddd;*/
    padding: 0.2rem 1rem;
    min-width: fit-content;
    white-space: nowrap;
    color: var( --ad-tableF);
    /* ไม่ตัดข้อความในเซลล์ */
    word-wrap: break-word;
    z-index: 0;
}
.Table:deep() td {
    padding: 0.75rem 1rem;
}

.Table:deep() th {
    background-color: none;
    position: sticky;
    /* ตั้งค่าให้ส่วนหัวตารางติดอยู่ด้านบนของตาราง */
    top: 0;
    /* ตั้งค่าให้ส่วนหัวตารางติดอยู่ที่ด้านบนของหน้าจอ */
    z-index: 5;
    background-color: var(--ad-tableHeadB)
}

/*tr:nth-child(even) {
    background-color: #f2f2f2;
}*/

.Table:deep() tr:nth-child(odd) {
    background-color: var(--ad-tableRowOddB);
}

.Table:deep() tr:nth-child(even) {
    background-color: var(--ad-tableRowEvenB);
}

.Table:deep() tr:hover {
    background-color: var(--ad-tableRow-H);
}

.ContentContainer {
    width: 90vw;
    background-color: var(--ad-conContainB);
}

.Table:deep() .UserHeader {
    text-align: left;
    position: sticky;
    padding-left: 1.5rem;
    left: 0;
    z-index: 4;
    background-color: var(--ad-UserHB);
    width: 90vw;
}

.AddGroupBtn {
    position: fixed;
    top: 94.5%;
    width: 2.3rem;
    height: 2.3rem;
    right: 2rem;
    border: 3px solid var(--ad-addGroupButB);
    padding: 0.3rem;
    border-radius: 50%;
    cursor: pointer;
}

.DropdownContainer {
    width: 10rem;
    height: fit-content;
    background-color: var(--ad-DropDownConB);
    position: fixed;
    bottom: 2.8rem;
    right: 2rem;
    list-style: none;
}

.DropdownMenu {
    padding: 1rem;
    font-size: 1rem;
    cursor: pointer;
}

.DropdownMenu:hover {
    background-color: var(--ad-DropDownmB-H);
    color: var(--ad-DropDownmF-H)
}
.controller-button {
  padding: 0.5rem 2rem;
  border: none;
  background: var(--editor-conMainB);
  color: var(--nav-signin);
  border-radius: 1rem;
  box-shadow: 0 0 0.25rem 0.1rem var(--editor_conShadow);
  transition: all .25s ease-in-out;
  cursor: pointer;
  font-weight: bold;
  /* margin: 0 0.4rem 0 1rem; */
  /* width: 8rem; */
}

.controller-button.submiting {
  cursor: not-allowed;
}

.controller-button:hover {
  background: var(--nav-F);
  color: var(--nav-btn-H);
  box-shadow: 0 0 0.25rem 0.1rem #ffffffab;
}
.exportwrapper{
    display: flex;
    justify-content: flex-end;
    background-color: var(--ad-menuconB);
    gap: 1rem;
}
</style>
