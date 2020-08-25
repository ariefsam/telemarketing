import Vue from 'vue'
import firebase from 'firebase'
import auth from 'firebase/auth' // not used but needed
import config from './environment.json'

export const fireApp = firebase.initializeApp(config)

fireApp.firestore().settings({
    cacheSizeBytes: firebase.firestore.CACHE_SIZE_UNLIMITED
});

fireApp.firestore().enablePersistence()
export const firestoreDB = fireApp.firestore();

export const AUTH = fireApp.auth()
export const provider = new firebase.auth.GoogleAuthProvider();

Vue.prototype.$auth = AUTH
Vue.prototype.$googleProvider = provider
Vue.prototype.$fireDB = firestoreDB
Vue.prototype.$fireApp = fireApp
