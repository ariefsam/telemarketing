<template>
  <q-page class="flex flex-center page-index">
    <!--<q-circular-progress
      indeterminate
      size="40px"
      :thickness="0.6"
      color="primary"
      center-color="grey-8"
      class="q-ma-md"
    />-->
    <q-spinner-pie
      color="primary"
      size="50px"
    />
  </q-page>
</template>

<script>
export default {
  name: 'Dashboard',

  mounted() {
    var vm = this
    var user = this.$authService.getUser();
    setTimeout(function(){
      if (user.IsAdmin) {
        vm.$router.replace({
          name: "admin-dashboard",
        });
      } else if (!user.IsAdmin) {
        vm.$router.replace({
          name: "telemarketer-dashboard",
        });
      }
    }, 1500);
  }
}
</script>
