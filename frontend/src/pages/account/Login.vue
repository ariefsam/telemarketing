<template>
  <q-layout>
    <q-page-container class="page-container-login">
      <q-page class="row justify-center items-center">
        <div class="col-md-5 login-container text-justify">
          <div class="title">Welcome,</div>
          <div class="subtitle">Please login by Google to access dashboard</div>
          <div
            class="row items-center justify-center social-button q-mt-md q-mb-sm"
            @click="loginByGoogle"
          >
            <img src="~assets/google-icon-small.png" width="22" />
            <div class="q-pl-sm">Login by Google</div>
          </div>
        </div>
      </q-page>
      <q-dialog v-model="unauthorizedDialog">
        <q-card>
          <q-card-section class="row items-center q-pb-none">
            <div class="text-h6 text-bold">Unauthorized Access</div>
            <q-space />
          </q-card-section>

          <q-card-section class="q-pt-sm" style="padding-bottom: 0">
            Access denied. Please log in with administrator privilege account.
          </q-card-section>

          <q-card-actions align="right">
            <q-btn flat label="OK" color="primary" v-close-popup />
          </q-card-actions>
        </q-card>
      </q-dialog>
    </q-page-container>
  </q-layout>
</template>

<script>
export default {
  name: "Login",

  data() {
    return {
      unauthorizedDialog: false,
    };
  },

  methods: {
    loginByGoogle() {
      var vm = this;
      vm.$auth
        .signInWithPopup(vm.$googleProvider)
        .then(function(result) {
          var token = result.credential.accessToken;
          var user = result.user;
          user.getIdToken(false).then(function(idToken) {
            var dataSubmit = {
              FirebaseToken: idToken
            };
            vm.$axios
              .post("/api/login/firebase", dataSubmit)
              .then(function(response) {
                if (response.data.Token) {
                  vm.$authService.setToken(response.data.Token);
                  vm.$router.replace({name : 'dashboard'});
                }
              })
              .catch(function(error) {
                vm.unauthorizedDialog = true
              });
          });
        })
        .catch(function(error) {
          var errorCode = error.code;
          var errorMessage = error.message;
          var email = error.email;
          var credential = error.credential;
          vm.unauthorizedDialog = true
        });
    },
  },
};
</script>