<template>
  <v-container fluid ma-0 pa-0>
    <Banner :img="'https://n-allo.be/wp-content/uploads/2016/08/ef3-placeholder-image-450x350.jpg'" />
    <br>
    <Loading :loading="loading" />
    <div v-if="!loading">
    <Slideshow v-if="info0.hasImg" :resource="'/bateauids/'" :infos="info0.arr" />
    <Slideshow v-if="info1.hasImg" :resource="'/comments/'" :infos="info1.arr" />
    <Slideshow v-if="info2.hasImg" :resource="'/interviews/'" :infos="info2.arr" />
    <Slideshow v-if="info3.hasImg" :resource="'/placeportids/'" :infos="info3.arr" />
    </div>
    <p class="footer">POC created by - TOOAP <img class="logo" src="./../assets/tooapSquare.png"></p>
  </v-container>
</template>


<script>
  import constants from "@/constants"
  import Slideshow from '@/components/Slideshow';
  import Banner from '@/components/Banner';
  import Loading from '@/components/Loading';
  const url0 = `${constants.apiUrl}bateauids`;
  const url1 = `${constants.apiUrl}comments`;
  const url2 = `${constants.apiUrl}interviews`;
  const url3 = `${constants.apiUrl}placeportids`;

  export default {
    components: {
      Slideshow,
      Banner,
      Loading
    },
    data: () => ({
      info0: { arr: [], hasImg: false },
      info1: { arr: [], hasImg: false },
      info2: { arr: [], hasImg: false },
      info3: { arr: [], hasImg: false },
      loading: true
    }),
    mounted() {
      this.$http.get(url0).then(res => {
        this.info0.arr = res.data;
        this.info0.arr = this.info0.arr.map(x => ({
          id: x.id || "", title: [x.nombateau_title,], img: []
        }))
        this.info0.hasImg = this.info0.arr.every(e => e.img.length > 0);
      }).catch(err => {
        throw err;
      });
      this.$http.get(url1).then(res => {
        this.info1.arr = res.data;
        this.info1.arr = this.info1.arr.map(x => ({
          id: x.id || "", title: [], img: []
        }))
        this.info1.hasImg = this.info1.arr.every(e => e.img.length > 0);
      }).catch(err => {
        throw err;
      });
      this.$http.get(url2).then(res => {
        this.info2.arr = res.data;
        this.info2.arr = this.info2.arr.map(x => ({
          id: x.id || "", title: [], img: []
        }))
        this.info2.hasImg = this.info2.arr.every(e => e.img.length > 0);
      }).catch(err => {
        throw err;
      });
      this.$http.get(url3).then(res => {
        this.info3.arr = res.data;
        this.info3.arr = this.info3.arr.map(x => ({
          id: x.id || "", title: [x.nomport_title,], img: [x.photo_img || "",]
        }))
        this.info3.hasImg = this.info3.arr.every(e => e.img.length > 0);
      }).catch(err => {
        throw err;
      });
      this.loading = false;
    }
  };
</script>

<style>
.footer {
  position: fixed;
  left: 0;
  bottom: 0;
  width: 100%;
  text-align: center;
}
.logo {
  width: 18px;
}
</style>