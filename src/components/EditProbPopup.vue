<script setup lang="ts" >
import { reactive } from 'vue';
import { useRouter } from 'vue-router';

const emitonEditProbPopupBtnClick = defineEmits(["onEditProbPopupBtnClick"])
const router = useRouter();
const DefineProbs = defineProps(["Probs"])

const state = reactive({
    Idfocus: false,
    Namefocus: false,
    Infofocus: false,
    Linkfocus: false,
})
const onUpdateBtnClick = async (e: Event) => {
    e.preventDefault();
    console.log(
            DefineProbs.Probs
        )
        try {
        const token = localStorage.getItem('token');
        if (!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api + "/api/sentry/updateassignment", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(
                {
                    Id:DefineProbs.Probs.Id,
                    Name:DefineProbs.Probs.Name,
                    Info:DefineProbs.Probs.Info,
                    Link:DefineProbs.Probs.Link
                }
            )
        });
        const json = await response.json();
        if (json.status === "success" && json.response) {
            emitonEditProbPopupBtnClick("onEditProbPopupBtnClick")
                router.go(0)
        } else {
            console.log("Updating Error", json)
        }
    } catch (error) {
        // data.status = "error";
    }

}
// }
</script>

<template>
    <div class="Container">
        <div class="icon-wrap">
            <div class="exit-button" @click="$emit('onEditProbPopupBtnClick')">
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
                Edit Assignment
            </div>
            <div class="controller">
                <form class="formWrap" @submit="onUpdateBtnClick">
                    <div :class="`inputWrapper ${state.Idfocus || (DefineProbs.Probs.Id !== '') ? 'focus' : ''}`">
                        <label class="placeholder">Proposition Id</label>
                        <input v-model="DefineProbs.Probs.Id" type="text" class="inputbox" @focus="() => state.Idfocus = true"
                            @blur="() => state.Idfocus = false" disabled />
                    </div>
                    <div :class="`inputWrapper ${state.Namefocus || (DefineProbs.Probs.Name !== '') ? 'focus' : ''}`">
                        <label class="placeholder">Proposition name</label>
                        <input v-model="DefineProbs.Probs.Name" type="text" class="inputbox" @focus="() => state.Namefocus = true"
                            @blur="() => state.Namefocus = false" />
                    </div>
                    <div :class="`inputWrapper ${state.Infofocus || (DefineProbs.Probs.Info !== '') ? 'focus' : ''}`">
                        <label class="placeholder">Proposition Info</label>
                        <input v-model="DefineProbs.Probs.Info" type="text" class="inputbox" @focus="() => state.Infofocus = true"
                            @blur="() => state.Infofocus = false" />
                    </div>
                    <div :class="`inputWrapper ${state.Linkfocus || (DefineProbs.Probs.Link !== '') ? 'focus' : ''}`">
                        <label class="placeholder">Proposition Link</label>
                        <input v-model="DefineProbs.Probs.Link" type="text" class="inputbox" @focus="() => state.Linkfocus = true"
                            @blur="() => state.Linkfocus = false" />
                    </div>
                    <button class="FormButton" type="submit">Update</button>

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