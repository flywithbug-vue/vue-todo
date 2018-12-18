import firebase from 'firebase'
import 'firebase/firestore'

const config = {
  apiKey: "AIzaSyCPGb3KatIhBtRzR6LtX8mjZd02iMYnx4k",
  authDomain: "todo-238f7.firebaseapp.com",
  databaseURL: "https://todo-238f7.firebaseio.com",
  projectId: "todo-238f7",
  storageBucket: "todo-238f7.appspot.com",
  messagingSenderId: "950396632840"
}

const firebaseApp = firebase.initializeApp(config)
const firestore = firebaseApp.firestore()
firestore.settings({timestampsInSnapshots: true})

export default firestore
