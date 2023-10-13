<script setup lang="ts">
import useUserStore from "../module/userstore"
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router'
import { RouterLink } from 'vue-router';
import useTheme from "../module/usetheme"
const themedata = useTheme();
const DropDownRef = ref()
const ButtonRef = ref()

const Dropdown = reactive<{ IsShow: boolean }>({
  IsShow: false
})
const DropdownShow = () => {
  Dropdown.IsShow = true;
  window.addEventListener("click", outSideDetected)
}
const userdata = useUserStore();
const router = useRouter();

const Logout = () => {
  localStorage.removeItem("token");
  router.push("/login");
}

const outSideDetected = (e: Event) => {
  try {
    if (!DropDownRef.value) {
      Dropdown.IsShow = false;
      window.removeEventListener("click", outSideDetected)
    }
    if (!DropDownRef.value.contains(e.target) && !ButtonRef.value.contains(e.target)) {
      Dropdown.IsShow = false;
      window.removeEventListener("click", outSideDetected)
    }
  } catch {
    window.removeEventListener("click", outSideDetected)
  }

}
</script>

<template>
  <div class="navigation">
    <div class="controller-schemewrap">
      <button class="scheme-selector" v-if="themedata.theme === 'light'" @click="themedata.setTheme('dark')">
        <svg viewBox="0 0 302.4 302.4">
          <g>
            <path
              d="M204.8 97.6C191.2 84 172 75.2 151.2 75.2s-40 8.4-53.6 22.4c-13.6 13.6-22.4 32.8-22.4 53.6s8.8 40 22.4 53.6c13.6 13.6 32.8 22.4 53.6 22.4s40-8.4 53.6-22.4c13.6-13.6 22.4-32.8 22.4-53.6s-8.4-40-22.4-53.6zm-14.4 92.8c-10 10-24 16-39.2 16s-29.2-6-39.2-16-16-24-16-39.2 6-29.2 16-39.2 24-16 39.2-16 29.2 6 39.2 16 16 24 16 39.2-6 29.2-16 39.2zM292 140.8h-30.8c-5.6 0-10.4 4.8-10.4 10.4 0 5.6 4.8 10.4 10.4 10.4H292c5.6 0 10.4-4.8 10.4-10.4 0-5.6-4.8-10.4-10.4-10.4zM151.2 250.8c-5.6 0-10.4 4.8-10.4 10.4V292c0 5.6 4.8 10.4 10.4 10.4 5.6 0 10.4-4.8 10.4-10.4v-30.8c0-5.6-4.8-10.4-10.4-10.4zM258 243.6l-22-22c-3.6-4-10.4-4-14.4 0s-4 10.4 0 14.4l22 22c4 4 10.4 4 14.4 0s4-10.4 0-14.4zM151.2 0c-5.6 0-10.4 4.8-10.4 10.4v30.8c0 5.6 4.8 10.4 10.4 10.4 5.6 0 10.4-4.8 10.4-10.4V10.4c0-5.6-4.8-10.4-10.4-10.4zM258.4 44.4c-4-4-10.4-4-14.4 0l-22 22c-4 4-4 10.4 0 14.4 3.6 4 10.4 4 14.4 0l22-22c4-4 4-10.4 0-14.4zM41.2 140.8H10.4c-5.6 0-10.4 4.8-10.4 10.4s4.4 10.4 10.4 10.4h30.8c5.6 0 10.4-4.8 10.4-10.4 0-5.6-4.8-10.4-10.4-10.4zM80.4 221.6c-3.6-4-10.4-4-14.4 0l-22 22c-4 4-4 10.4 0 14.4s10.4 4 14.4 0l22-22c4-4 4-10.4 0-14.4zM80.4 66.4l-22-22c-4-4-10.4-4-14.4 0s-4 10.4 0 14.4l22 22c4 4 10.4 4 14.4 0s4-10.4 0-14.4z">
            </path>
          </g>
        </svg>
      </button>
      <button class="scheme-selector" v-else @click="themedata.setTheme('light')">
        <svg viewBox="0 0 312.999 312.999">
          <g>
            <path
              d="M305.6 178.053c-3.2-.8-6.4 0-9.2 2-10.4 8.8-22.4 16-35.6 20.8-12.4 4.8-26 7.2-40.4 7.2-32.4 0-62-13.2-83.2-34.4-21.2-21.2-34.4-50.8-34.4-83.2 0-13.6 2.4-26.8 6.4-38.8 4.4-12.8 10.8-24.4 19.2-34.4 3.6-4.4 2.8-10.8-1.6-14.4-2.8-2-6-2.8-9.2-2-34 9.2-63.6 29.6-84.8 56.8-20.4 26.8-32.8 60-32.8 96.4 0 43.6 17.6 83.2 46.4 112s68.4 46.4 112 46.4c36.8 0 70.8-12.8 98-34 27.6-21.6 47.6-52.4 56-87.6 2-6-1.2-11.6-6.8-12.8zm-61.2 83.6c-23.2 18.4-52.8 29.6-85.2 29.6-38 0-72.4-15.6-97.2-40.4-24.8-24.8-40.4-59.2-40.4-97.2 0-31.6 10.4-60.4 28.4-83.6 12.4-16 28-29.2 46-38.4-2 4.4-4 8.8-5.6 13.6-5.2 14.4-7.6 29.6-7.6 45.6 0 38 15.6 72.8 40.4 97.6s59.6 40.4 97.6 40.4c16.8 0 32.8-2.8 47.6-8.4 5.2-2 10.4-4 15.2-6.4-9.6 18.4-22.8 34.8-39.2 47.6z">
            </path>
          </g>
        </svg>
      </button>
    </div>
    <div class="imgwrapper">

      <svg viewBox="0 0 800 800">
        <defs id="defs6">
          <linearGradient x1="0" y1="0" x2="1" y2="0" gradientUnits="userSpaceOnUse"
            gradientTransform="matrix(-75.98391,162.94803,162.94803,75.98391,306.65924,276.96521)" spreadMethod="pad"
            id="linearGradient30">
            <stop style="stop-opacity:1;stop-color:#c8e9f1" offset="0" id="stop22" />
            <stop style="stop-opacity:1;stop-color:#c8e9f1" offset="0.2" id="stop24" />
            <stop style="stop-opacity:1;stop-color:#008fca" offset="0.8" id="stop26" />
            <stop style="stop-opacity:1;stop-color:#008fca" offset="1" id="stop28" />
          </linearGradient>
          <linearGradient x1="0" y1="0" x2="1" y2="0" gradientUnits="userSpaceOnUse"
            gradientTransform="matrix(119.97658,207.80553,207.80553,-119.97658,210.49591,238.09319)" spreadMethod="pad"
            id="linearGradient50">
            <stop style="stop-opacity:1;stop-color:#7ba0ac" offset="0" id="stop46" />
            <stop style="stop-opacity:1;stop-color:#95c7eb" offset="1" id="stop48" />
          </linearGradient>
          <linearGradient x1="0" y1="0" x2="1" y2="0" gradientUnits="userSpaceOnUse"
            gradientTransform="matrix(0,332.29007,332.29007,0,209.36749,173.44919)" spreadMethod="pad"
            id="linearGradient72">
            <stop style="stop-opacity:1;stop-color:#eb3e27" offset="0" id="stop66" />
            <stop style="stop-opacity:1;stop-color:#f8dd18" offset="0.8" id="stop68" />
            <stop style="stop-opacity:1;stop-color:#f8dd18" offset="1" id="stop70" />
          </linearGradient>
          <clipPath clipPathUnits="userSpaceOnUse" id="clipPath82">
            <path d="M 0,600 H 600 V 0 H 0 Z" id="path80" />
          </clipPath>
          <linearGradient x1="0" y1="0" x2="1" y2="0" gradientUnits="userSpaceOnUse"
            gradientTransform="matrix(0,342.221,-342.221,0,-119.7873,453.0068)" spreadMethod="pad" id="linearGradient94">
            <stop style="stop-opacity:1;stop-color:#fff33b" offset="0" id="stop86" />
            <stop style="stop-opacity:1;stop-color:#fff33b" offset="0.47465006" id="stop88" />
            <stop style="stop-opacity:1;stop-color:#e93e39" offset="0.7" id="stop90" />
            <stop style="stop-opacity:1;stop-color:#e93e39" offset="1" id="stop92" />
          </linearGradient>
        </defs>
        <g id="g8" inkscape:groupmode="layer" inkscape:label="thnicf-academy"
          transform="matrix(1.3333333,0,0,-1.3333333,0,800)">
          <g id="g10">
            <g id="g12">
              <g id="g18">
                <g id="g20">
                  <path
                    d="m 146.689,328.639 c -0.908,-17.259 6.365,-43.37 18.439,-59.164 v 0 c 26.461,-29.411 54.548,-26.456 55.264,-26.456 v 0 c 33.468,0 60.599,27.129 60.599,60.599 v 0 c 0,32.817 -26.094,59.525 -58.662,60.55 v 0 c 20.471,30.197 55.064,50.043 94.294,50.043 v 0 c 34.595,-0.778 78.48,-17.155 97.151,-54.516 v 0 c -18.37,56.366 -66.607,88.596 -129.102,88.596 v 0 c -77.733,0 -138.059,-52.327 -137.983,-119.652"
                    style="fill:url(#linearGradient30);stroke:none" id="path32" />
                </g>
              </g>
            </g>
          </g>
          <g id="g34">
            <g id="g36">
              <g id="g42">
                <g id="g44">
                  <path
                    d="m 146.689,328.639 c -0.908,-17.259 6.365,-43.37 18.439,-59.164 v 0 c 22.033,-24.489 45.194,-26.539 52.742,-26.539 v 0 c 1.518,0 2.402,0.083 2.522,0.083 v 0 c 33.468,0 60.599,27.129 60.599,60.599 v 0 c 0,32.817 -26.094,59.525 -58.662,60.55 v 0 c 20.471,30.197 55.064,50.043 94.294,50.043 v 0 c 34.595,-0.778 78.48,-17.155 97.151,-54.516 v 0 c -18.37,56.366 -66.607,88.596 -129.102,88.596 v 0 c -77.733,0 -138.059,-52.327 -137.983,-119.652 m 51.652,-77.874 c -10.995,4.142 -21.132,11.296 -30.131,21.267 v 0 c -12,15.787 -18.335,41.044 -17.527,56.397 v 0 l 0.006,0.107 v 0.107 c -0.016,14.433 2.834,28.41 8.471,41.543 v 0 c 5.47,12.741 13.42,24.383 23.632,34.606 v 0 c 12.311,12.326 27.281,22.033 44.493,28.854 v 0 c 17.823,7.063 37.13,10.644 57.387,10.644 v 0 c 30.009,0 57.273,-7.768 78.845,-22.465 v 0 c 12.479,-8.503 23.095,-19.364 31.516,-32.131 v 0 c -6.494,5.554 -13.886,10.44 -22.066,14.552 v 0 c -17.056,8.572 -37.033,13.532 -56.254,13.964 v 0 l -0.045,0.001 h -0.045 c -39.13,0 -75.617,-19.364 -97.605,-51.798 v 0 l -4.077,-6.015 7.262,-0.229 c 14.735,-0.463 28.514,-6.551 38.801,-17.141 v 0 c 10.309,-10.614 15.987,-24.61 15.987,-39.41 v 0 c 0,-31.209 -25.391,-56.6 -56.599,-56.6 v 0 c -0.061,0 -0.21,-0.002 -0.439,-0.017 v 0 c -0.333,-0.021 -1.024,-0.064 -2.083,-0.064 v 0 c -3.496,0 -10.683,0.496 -19.529,3.828"
                    style="fill:url(#linearGradient50);stroke:none" id="path52" />
                </g>
              </g>
            </g>
          </g>
          <g id="g54">
            <g id="g56">
              <g id="g62">
                <g id="g64">
                  <path
                    d="m 59.503,343.459 c 0,-93.896 76.115,-170.01 170.009,-170.01 v 0 c 16.515,0 32.48,2.354 47.578,6.745 v 0 c 30.907,8.992 51.725,26.421 65.372,49.65 v 0 c 7.91,13.465 15.956,35.093 16.713,59.186 v 0 c 0.793,25.194 -6.384,53.086 -30.021,77.147 v 0 c -3.974,4.5 -8.497,8.504 -13.483,11.882 v 0 c -0.006,0.005 -0.013,0.009 -0.019,0.015 v 0 c -11.377,7.703 -25.101,12.204 -39.875,12.204 v 0 c -21.73,0 -35.617,-10.003 -48.678,-25.349 v 0 c 36.617,-2.611 56.93,-32.639 56.627,-62.335 v 0 c 0,-21.653 -10.938,-40.735 -27.558,-51.976 v 0 c 0.041,-0.009 0.083,-0.02 0.123,-0.031 v 0 c -11.341,-6.82 -23.775,-10.59 -36.819,-10.59 v 0 c -53.12,0 -96.185,62.425 -96.185,139.434 v 0 c 0,55.873 22.673,104.062 55.406,126.308 v 0 C 109.624,484.133 59.503,419.65 59.503,343.459"
                    style="fill:url(#linearGradient72);stroke:none" id="path74" />
                </g>
              </g>
            </g>
          </g>
          <g id="g76">
            <g id="g78" clip-path="url(#clipPath82)">
              <g id="g84" transform="translate(329.1543,366.1768)">
                <path
                  d="m 0,0 c -3.975,4.5 -8.497,8.504 -13.483,11.882 -0.006,0.005 -0.013,0.01 -0.02,0.015 -11.377,7.703 -25.1,12.204 -39.874,12.204 -21.731,0 -35.618,-10.003 -48.679,-25.348 36.618,-2.611 56.931,-32.64 56.628,-62.336 0,-21.652 -10.938,-40.735 -27.558,-51.976 0.041,-0.009 0.083,-0.02 0.123,-0.031 -11.341,-6.82 -23.776,-10.59 -36.819,-10.59 -53.121,0 -96.185,62.426 -96.185,139.434 0,55.873 22.673,104.061 55.406,126.308 -69.069,-21.606 -119.191,-86.089 -119.191,-162.28 0,-93.896 76.116,-170.01 170.009,-170.01 16.515,0 32.481,2.355 47.579,6.746 30.907,8.992 51.725,26.42 65.372,49.649 7.91,13.465 15.956,35.094 16.713,59.186 C 30.813,-51.953 23.637,-24.061 0,0 Z"
                  style="fill:none;stroke:url(#linearGradient94);stroke-width:5;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:10;stroke-dasharray:none;stroke-opacity:1"
                  id="path96" />
              </g>
              <g id="g98" transform="translate(238.2417,322.7935)">
                <path
                  d="M 0,0 -3.498,-18.832 H -24.96 l -14.467,-78.793 h -18.281 l 14.466,78.793 H -65.498 L -62.001,0 Z"
                  style="fill:#254090;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path100" />
              </g>
              <g id="g102" transform="translate(238.2417,322.7935)">
                <path
                  d="M 0,0 -3.498,-18.832 H -24.96 l -14.467,-78.793 h -18.281 l 14.466,78.793 H -65.498 L -62.001,0 Z"
                  style="fill:none;stroke:#ffffff;stroke-width:2;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:10;stroke-dasharray:none;stroke-opacity:1"
                  id="path104" />
              </g>
              <g id="g106" transform="translate(318.918,322.791)">
                <path
                  d="m 0,0 -4.187,-19.114 -42.91,0.075 L -44.089,0 h -18.282 l -3.504,-19.039 h -9.688 l -3.975,-17.579 h 10.429 l -11.225,-61.005 h 18.122 l 11.605,61.005 h 42.923 l -10.12,-61.005 H 0.478 L 18.441,0 Z"
                  style="fill:#254090;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path108" />
              </g>
              <g id="g110" transform="translate(318.918,322.791)">
                <path
                  d="m 0,0 -4.187,-19.114 -42.91,0.075 L -44.089,0 h -18.282 l -3.504,-19.039 h -9.688 l -3.975,-17.579 h 10.429 l -11.225,-61.005 h 18.122 l 11.605,61.005 h 42.923 l -10.12,-61.005 H 0.478 L 18.441,0 Z"
                  style="fill:none;stroke:#ffffff;stroke-width:2;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:10;stroke-dasharray:none;stroke-opacity:1"
                  id="path112" />
              </g>
              <g id="g114" transform="translate(437.1807,322.791)">
                <path
                  d="m 0,0 -17.965,-97.623 h -18.76 l 11.1,62.242 c 0,0 -8.281,13.431 -44.08,20.278 l -15.559,-82.52 h -18.282 L -85.582,0 h 18.6 c 33.407,-5.39 45.334,-17.365 45.334,-17.365 L -18.441,0 Z"
                  style="fill:#234390;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path116" />
              </g>
              <g id="g118" transform="translate(437.1807,322.791)">
                <path
                  d="m 0,0 -17.965,-97.623 h -18.76 l 11.1,62.242 c 0,0 -8.281,13.431 -44.08,20.278 l -15.559,-82.52 h -18.282 L -85.582,0 h 18.6 c 33.407,-5.39 45.334,-17.365 45.334,-17.365 L -18.441,0 Z"
                  style="fill:none;stroke:#ffffff;stroke-width:2;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:10;stroke-dasharray:none;stroke-opacity:1"
                  id="path120" />
              </g>
              <g id="g122" transform="translate(466.7529,303.2065)">
                <path
                  d="m 0,0 c 0,-4.937 -4.002,-8.937 -8.938,-8.937 -4.937,0 -8.939,4 -8.939,8.937 0,4.937 4.002,8.939 8.939,8.939 C -4.002,8.939 0,4.937 0,0 m -31.64,-78.039 12.346,64.384 h 17.489 l -12.347,-64.384 z"
                  style="fill:#254090;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path124" />
              </g>
              <g id="g126" transform="translate(466.7529,303.2065)">
                <path
                  d="m 0,0 c 0,-4.937 -4.002,-8.937 -8.938,-8.937 -4.937,0 -8.939,4 -8.939,8.937 0,4.937 4.002,8.939 8.939,8.939 C -4.002,8.939 0,4.937 0,0 Z m -31.64,-78.039 12.346,64.384 h 17.489 l -12.347,-64.384 z"
                  style="fill:none;stroke:#ffffff;stroke-width:2;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:10;stroke-dasharray:none;stroke-opacity:1"
                  id="path128" />
              </g>
              <g id="g130" transform="translate(523.7812,225.6465)">
                <path
                  d="m 0,0 c 0,0 -5.512,-1.322 -8.425,-1.754 -2.916,-0.43 -5.804,-0.647 -8.664,-0.647 -6.042,0 -11.343,1.281 -15.897,3.839 -4.56,2.56 -8.243,6.063 -11.049,10.513 -2.81,4.448 -4.665,9.691 -5.566,15.722 -0.9,6.034 -0.714,12.523 0.558,19.472 1.375,7.678 3.788,14.716 7.233,21.116 3.443,6.398 7.496,11.884 12.162,16.453 4.662,4.57 9.8,8.104 15.419,10.604 5.618,2.497 11.34,3.749 17.17,3.749 6.676,0 11.472,-0.779 11.472,-0.779 L 10.175,78.59 c 0,0 -7.499,1.281 -10.891,1.281 -5.192,0 -9.593,-1.4 -13.194,-4.195 -3.605,-2.797 -6.597,-6.018 -8.982,-9.664 -2.385,-3.646 -4.189,-7.289 -5.405,-10.935 -1.22,-3.645 -1.988,-6.259 -2.306,-7.838 -0.849,-4.496 -1.06,-8.631 -0.634,-12.398 0.422,-3.768 1.429,-7.018 3.019,-9.753 1.59,-2.734 3.709,-4.861 6.359,-6.38 2.648,-1.52 5.775,-2.277 9.38,-2.277 4.026,0 12.002,2.285 12.002,2.285 z"
                  style="fill:#254090;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path132" />
              </g>
              <g id="g134" transform="translate(523.7812,225.6465)">
                <path
                  d="m 0,0 c 0,0 -5.512,-1.322 -8.425,-1.754 -2.916,-0.43 -5.804,-0.647 -8.664,-0.647 -6.042,0 -11.343,1.281 -15.897,3.839 -4.56,2.56 -8.243,6.063 -11.049,10.513 -2.81,4.448 -4.665,9.691 -5.566,15.722 -0.9,6.034 -0.714,12.523 0.558,19.472 1.375,7.678 3.788,14.716 7.233,21.116 3.443,6.398 7.496,11.884 12.162,16.453 4.662,4.57 9.8,8.104 15.419,10.604 5.618,2.497 11.34,3.749 17.17,3.749 6.676,0 11.472,-0.779 11.472,-0.779 L 10.175,78.59 c 0,0 -7.499,1.281 -10.891,1.281 -5.192,0 -9.593,-1.4 -13.194,-4.195 -3.605,-2.797 -6.597,-6.018 -8.982,-9.664 -2.385,-3.646 -4.189,-7.289 -5.405,-10.935 -1.22,-3.645 -1.988,-6.259 -2.306,-7.838 -0.849,-4.496 -1.06,-8.631 -0.634,-12.398 0.422,-3.768 1.429,-7.018 3.019,-9.753 1.59,-2.734 3.709,-4.861 6.359,-6.38 2.648,-1.52 5.775,-2.277 9.38,-2.277 4.026,0 12.002,2.285 12.002,2.285 z"
                  style="fill:none;stroke:#ffffff;stroke-width:2;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:10;stroke-dasharray:none;stroke-opacity:1"
                  id="path136" />
              </g>
              <g id="g138" transform="translate(200.8345,181.042)">
                <path
                  d="M 0,0 -1.913,5.346 -3.826,0 Z M 6.173,-17.22 3.429,-9.625 H -7.255 l -2.744,-7.595 h -11.092 l 14.529,39.954 h 9.35 L 17.259,-17.22 Z"
                  style="fill:#ffffff;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path140" />
              </g>
              <g id="g142" transform="translate(234.103,163.5283)">
                <path
                  d="m 0,0 c -11.431,0 -20.051,8.768 -20.051,20.394 0,11.681 8.62,20.491 20.051,20.491 8.706,0 15.669,-4.583 18.626,-12.26 l 1.57,-4.078 H 8.55 L 7.752,26.311 C 6.354,29.399 3.746,30.966 0,30.966 c -5.768,0 -9.494,-4.15 -9.494,-10.572 0,-6.393 3.726,-10.523 9.494,-10.523 3.744,0 6.354,1.554 7.758,4.618 L 8.56,16.24 H 20.208 L 18.623,12.155 C 15.669,4.544 8.707,0 0,0"
                  style="fill:#ffffff;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path144" />
              </g>
              <g id="g146" transform="translate(270.7075,181.042)">
                <path
                  d="M 0,0 -1.913,5.346 -3.826,0 Z M 6.173,-17.22 3.429,-9.625 H -7.255 l -2.744,-7.595 h -11.092 l 14.529,39.954 h 9.35 L 17.259,-17.22 Z"
                  style="fill:#ffffff;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path148" />
              </g>
              <g id="g150" transform="translate(299.2241,173.4971)">
                <path
                  d="M 0,0 C 7.014,0 10.571,3.475 10.571,10.327 10.571,17.31 7.014,20.85 0,20.85 H -3.174 V 0 Z M -13.633,-9.675 V 30.476 H 0 c 12.834,0 21.127,-7.91 21.127,-20.149 C 21.127,-1.823 12.834,-9.675 0,-9.675 Z"
                  style="fill:#ffffff;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path152" />
              </g>
              <g id="g154" transform="translate(319.7432,163.8223)">
                <path
                  d="M 0,0 V 40.199 H 24.374 V 30.524 H 10.459 V 25.108 H 22.904 V 15.434 H 10.459 V 9.675 H 24.374 V 0 Z"
                  style="fill:#ffffff;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path156" />
              </g>
              <g id="g158" transform="translate(374.5713,163.8223)">
                <path
                  d="M 0,0 V 14.183 L -6.321,0 h -7.029 l -6.373,14.263 V 0 h -10.459 v 39.905 h 9.746 L -9.837,16.224 0.762,39.905 h 9.697 V 0 Z"
                  style="fill:#ffffff;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path160" />
              </g>
              <g id="g162" transform="translate(394.417,163.8223)">
                <path
                  d="M 0,0 V 15.146 L -13.136,40.15 H -1.347 L 5.205,26.704 11.758,40.15 H 23.536 L 10.459,15.148 V 0 Z"
                  style="fill:#ffffff;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path164" />
              </g>
              <g id="g166" transform="translate(198.9214,195.2891)">
                <path
                  d="M 0,0 -6.174,-17.247 H 6.174 Z M 7.447,-20.872 H -7.447 l -2.744,-7.595 h -4.704 L -2.548,5.487 h 5.145 l 12.298,-33.954 h -4.704 z"
                  style="fill:#008fca;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path168" />
              </g>
              <g id="g170" transform="translate(234.103,201.4131)">
                <path
                  d="M 0,0 C 7.398,0 13.278,-3.724 15.826,-10.338 H 10.485 C 8.623,-6.222 4.998,-3.919 0,-3.919 c -7.153,0 -12.494,-5.194 -12.494,-13.572 0,-8.329 5.341,-13.523 12.494,-13.523 4.998,0 8.623,2.303 10.485,6.369 h 5.341 C 13.278,-31.21 7.398,-34.885 0,-34.885 c -9.554,0 -17.051,7.154 -17.051,17.394 C -17.051,-7.251 -9.554,0 0,0"
                  style="fill:#008fca;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path172" />
              </g>
              <g id="g174" transform="translate(268.7944,195.2891)">
                <path
                  d="M 0,0 -6.174,-17.247 H 6.174 Z M 7.447,-20.872 H -7.447 l -2.744,-7.595 h -4.704 L -2.548,5.487 h 5.145 l 12.298,-33.954 h -4.704 z"
                  style="fill:#008fca;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path176" />
              </g>
              <g id="g178" transform="translate(299.2241,170.4971)">
                <path
                  d="M 0,0 C 8.917,0 13.571,4.997 13.571,13.327 13.571,21.656 8.917,26.85 0,26.85 H -6.174 V 0 Z M 18.127,13.327 C 18.127,2.744 11.219,-3.675 0,-3.675 H -10.633 V 30.476 H 0 c 11.219,0 18.127,-6.566 18.127,-17.149"
                  style="fill:#008fca;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path180" />
              </g>
              <g id="g182" transform="translate(341.1172,197.3467)">
                <path
                  d="M 0,0 H -13.915 V -11.416 H -1.47 v -3.675 H -13.915 V -26.85 H 0 v -3.674 H -18.374 V 3.675 L 0,3.675 Z"
                  style="fill:#008fca;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path184" />
              </g>
              <g id="g186" transform="translate(347.3896,200.7275)">
                <path
                  d="M 0,0 H 4.802 L 17.345,-28.025 29.888,0 h 4.753 V -33.905 H 30.182 V -8.623 L 18.913,-33.905 H 15.777 L 4.459,-8.574 V -33.905 H 0 Z"
                  style="fill:#008fca;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path188" />
              </g>
              <g id="g190" transform="translate(386.2461,200.9727)">
                <path d="M 0,0 H 4.949 L 13.376,-17.296 21.804,0 h 4.948 L 15.63,-21.265 V -34.15 h -4.459 v 12.885 z"
                  style="fill:#008fca;fill-opacity:1;fill-rule:nonzero;stroke:none" id="path192" />
              </g>
            </g>
          </g>
        </g>
      </svg>

    </div>
    <div class="controller-accountwrap">
      <div class="UserName">
        {{ userdata.info?.firstname }}
      </div>
      <div class="accoute-manage-wrap" v-if="userdata.info">
        <button ref="ButtonRef" class="account-manage" @click="DropdownShow">
          <svg xmlns="http://www.w3.org/2000/svg" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink"
            xmlns:svgjs="http://svgjs.com/svgjs" x="0" y="0" viewBox="0 0 512 512" xml:space="preserve">
            <g>
              <path
                d="M437.02 74.98C388.668 26.63 324.379 0 256 0S123.332 26.629 74.98 74.98C26.63 123.332 0 187.621 0 256s26.629 132.668 74.98 181.02C123.332 485.37 187.621 512 256 512s132.668-26.629 181.02-74.98C485.37 388.668 512 324.379 512 256s-26.629-132.668-74.98-181.02zM111.105 429.297c8.454-72.735 70.989-128.89 144.895-128.89 38.96 0 75.598 15.179 103.156 42.734 23.281 23.285 37.965 53.687 41.742 86.152C361.641 462.172 311.094 482 256 482s-105.637-19.824-144.895-52.703zM256 269.507c-42.871 0-77.754-34.882-77.754-77.753C178.246 148.879 213.13 114 256 114s77.754 34.879 77.754 77.754c0 42.871-34.883 77.754-77.754 77.754zm170.719 134.427a175.9 175.9 0 0 0-46.352-82.004c-18.437-18.438-40.25-32.27-64.039-40.938 28.598-19.394 47.426-52.16 47.426-89.238C363.754 132.34 315.414 84 256 84s-107.754 48.34-107.754 107.754c0 37.098 18.844 69.875 47.465 89.266-21.887 7.976-42.14 20.308-59.566 36.542-25.235 23.5-42.758 53.465-50.883 86.348C50.852 364.242 30 312.512 30 256 30 131.383 131.383 30 256 30s226 101.383 226 226c0 56.523-20.86 108.266-55.281 147.934zm0 0"
                fill="#000000" data-original="#000000" class=""></path>
            </g>
          </svg>
        </button>
      </div>
      <div v-else style="display: flex;height: 100%;align-items: center;">
        <router-link to="/login" class="controller-button">
          <div class="signin">Sign in</div>
        </router-link>
      </div>
    </div>
    <ul ref="DropDownRef" class="DropdownContainer" v-if="Dropdown.IsShow">
      <router-link to="/admin">
        <li class="DropdownMenu">Admin</li>

      </router-link>
      <router-link to="/editor">
        <li class="DropdownMenu">
          Editor</li>
      </router-link>
      <li class="DropdownMenu" @click="Logout">Sign out</li>
    </ul>
  </div>
</template>

<style scoped>
.navigation {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  padding: .25rem 0;
  background: var(--nav-B);
  position: relative;
  overflow: hidden;
}

.imgwrapper {
  height: 4rem;
  aspect-ratio: 1/1;
}

.imgwrapper svg {
  width: 100%;
  height: 100%;
  transform: scale(1.5);
}


.account-manage {
  width: 100%;
  height: 100%;
  cursor: pointer;
  background: transparent;
  border: none;
  display: flex;
}

.account-manage path {
  fill: var(--nav-accButB);
}

.account-manage:hover path {
  fill: var(--nav-accButB-H);
}

.controller-accountwrap {
  right: 1%;
  position: absolute;
  height: 60%;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.accoute-manage-wrap {

  height: 100%;
  aspect-ratio: 1/1;
}

.UserName {
  font-size: 1.3rem;
  color: var(--nav-F);
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

.controller-button>div {
  text-align: center;
  width: 4rem;
}

.DropdownContainer {
  width: 10rem;
  height: fit-content;
  background-color: var(--nav-DropDownB);
  position: fixed;
  top: 5rem;
  right: 0;
  list-style: none;
  z-index: 10;
}

.DropdownMenu {
  padding: 1rem;
  font-size: 1rem;
  cursor: pointer;

}

.DropdownContainer>a {
  color: rgb(16, 16, 16);
  text-decoration: none;
}

.DropdownMenu:hover {
  background-color: var(--nav-DropDownB-H);
  color: var(--nav-DropDownF-H);
}

.controller-schemewrap {
  left: 1%;
  display: flex;
  align-items: center;
  position: absolute;
}

.scheme-selector {
  background: transparent;
  border: none;
  cursor: pointer;
  width: 2rem;
  aspect-ratio: 1/1;
}

.scheme-selector path {
  fill: var(--nav-schemeB);
}
</style>
