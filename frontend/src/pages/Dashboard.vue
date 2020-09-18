<template>
  <q-page class="flex flex-center">
    <q-circular-progress
      indeterminate
      size="40px"
      :thickness="0.6"
      color="primary"
      center-color="grey-8"
      class="q-ma-md"
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
      if (user.isAdmin) {
        vm.$router.replace({
          name: "admin-dashboard",
        });
      } else if (!user.isAdmin) {
        vm.$router.replace({
          name: "telemarketer-dashboard",
        });
      }
    }, 1000);
  }
}
</script>
