<script setup lang="ts">
import { useRoute } from "vue-router"
import { ref, watch } from "vue"
import api from "@/api"

const route = useRoute()
const pageData = ref(null)

const fetchData = async () => {
  const resp = await api.v1QuranPage(route.params.layoutId, route.params.pageNumber)
  console.log('api.v1QuranPage:', resp)
  if (resp !== null) {
    pageData.value = resp
  }
}

const nextPage = (params) => {
  let pageNumber = parseInt(params.pageNumber) + 1
  return {
    layoutId: params.layoutId,
    pageNumber,
  }
}

const prevPage = (params) => {
  let pageNumber = parseInt(params.pageNumber) - 1
  return {
    layoutId: params.layoutId,
    pageNumber,
  }
}

watch(route, () => fetchData(), { immediate: true })

</script>

<template>
  <div>
    <router-link :to="{ name: 'quranPage', params: nextPage($route.params) }">Next</router-link>
    Layout {{$route.params.layoutId}}, Page #{{$route.params.pageNumber}}
    <div style="text-align: justify; width: 800px;">
      <p v-for="line in pageData.listLine"  class="arab" style="font-size: 26px; padding-top: 10px; display: block; border-bottom: dotted 1px #00ff00; margin-bottom: 5px; direction: rtl">
        <template v-for="ayat in line.listAyat">
          {{ayat.text}} ({{ayat.ayatId}})
        </template>
      </p>
    </div>
    <router-link  :to="{ name: 'quranPage', params: prevPage($route.params) }">Prev</router-link>
  </div>
</template>