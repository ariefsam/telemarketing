<template>
  <q-layout>
    
    <q-page-container class="page-container">
      <q-page class="row justify-center ">
        <div class="col-md-5 login-container text-justify">
          
        </div>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
export default {
  name: "Login",

  data() {
    return {
      email: "",
    };
  },
  mounted() {
    var vm=this;
    var data_submit = {
      Token: vm.$authService.getToken(),
      Limit: 10000,
    }
    this.$axios
        .post("/api/customer", data_submit)
        .then(function (response) {
          if (response.data) {
            console.log(response.data);
          }
        })
  }
}
</script>