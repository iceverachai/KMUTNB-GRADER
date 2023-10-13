<script setup lang="ts" >
import { reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router';
const onEditUserClick = defineEmits(["onEditUserClick"])
const userName = defineProps(["userName"])
////////////////////////////////////////////
type UserDataType = {
    response: {
        
        data: {
            Username: string,
            Firstname: string,
            Lastname: string,
        }
    },
    status: string
}
////////////////////////////////////////////
const userData = reactive<UserDataType>({} as UserDataType)
const route = useRouter();

const state = reactive<{
    Usernamefocus: boolean,
    Passwordfocus: boolean,
    Firstnamefocus: boolean,
    Lastnamefocus: boolean,
}>({
    Usernamefocus: false,
    Passwordfocus: false,
    Firstnamefocus: false,
    Lastnamefocus: false,
})
//getUserDataToUpdate
// const GetUserData = async () => {
//     try {
//         const token = localStorage.getItem('token');
//         if (!token?.length) throw new Error('Token empty or undefined');

//         const response = await fetch(import.meta.env.VITE_api + "/api/sentry/getUserDataToUpdate", {
//             method: "GET",
//             headers: {
//                 "Content-Type": "application/json",
//                 "Authorization": `Bearer ${token}`
//             }
//         });
//         const json = await response.json();
//         if (json.status === "success" && json.response) {
//             userData.response.data = json.response.data;
//             userData.status = json.status;
//             return
//         } else {
//             userData.status = json.status;

//         }
//     } catch (error) {
//         userData.status = "error";
//     }
// }
onMounted(() => {
    // GetUserData()
})

const onUpdateBtnClick = async (e: Event) => {
    e.preventDefault();
    console.log(userName)
    try {
        // const token = localStorage.getItem('token');
        // if (!token?.length) throw new Error('Token empty or undefined');

        // const response = await fetch(import.meta.env.VITE_api + "/api/sentry/updateuser", {
        //     method: "POST",
        //     headers: {
        //         "Content-Type": "application/json",
        //         "Authorization": `Bearer ${token}`
        //     },
        //     body: JSON.stringify(
        //         {
        //             Username:userData.response.data.Username,
        //             Firstname:userData.response.data.Firstname,
        //             Lastname:userData.response.data.Lastname,
        //         }
        //     )
        // });
        // const json = await response.json();
        // if (json.status === "success" && json.response) {
        onEditUserClick("onEditUserClick")
        route.go(0)
        // } else {
        // console.log("Updating Error", json)
        // }
    } catch (error) {
        // data.status = "error";
    }
}
</script>

<template>
    <div class="Container">
        <div class="icon-wrap">
            <div class="exit-button" @click="$emit('onEditUserClick')">
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
                Edit user
            </div>
            <div class="controller">
                <form class="formWrap" @submit="onUpdateBtnClick">
                    <div :class="`inputWrapper ${state.Usernamefocus || (userData.response.data.Username !== '') ? 'focus' : ''}`">
                        <label class="placeholder">Username</label>
                        <input v-model="userData.response.data.Username" type="text" class="inputbox"
                            @focus="() => state.Usernamefocus = true" @blur="() => state.Usernamefocus = false" disabled />
                    </div>
                    <div :class="`inputWrapper ${state.Firstnamefocus || (userData.response.data.Firstname !== '') ? 'focus' : ''}`">
                        <label class="placeholder">Firstname</label>
                        <input v-model="userData.response.data.Firstname" type="text" class="inputbox"
                            @focus="() => state.Firstnamefocus = true" @blur="() => state.Firstnamefocus = false" />
                    </div>
                    <div :class="`inputWrapper ${state.Lastnamefocus || (userData.response.data.Lastname !== '') ? 'focus' : ''}`">
                        <label class="placeholder">Lastname</label>
                        <input v-model="userData.response.data.Lastname" type="text" class="inputbox"
                            @focus="() => state.Lastnamefocus = true" @blur="() => state.Lastnamefocus = false" />
                    </div>
                    <button class="FormButton" type="submit">Create</button>
                </form>
                <!-- <div v-if="fallback.length" class="fallbackerr">
                        {fallback}
                    </div> -->
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
}

.icon-wrap {
    padding: 2rem;
    background-color: var(--ad-modalIconWrapB);
    border-radius: 2rem;
    font-size: 1.2rem;
    position: relative;
    text-align: center;
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


</style>