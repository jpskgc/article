import * as firebase from 'firebase';
import 'firebase/auth';

const firebaseConfig = {
  apiKey: 'AIzaSyAaUYZ1icS_t_7OjH9f-VpDH-WGEIon_hs',
  authDomain: 'jpskgc-article.firebaseapp.com',
  databaseURL: 'https://jpskgc-article.firebaseio.com',
  projectId: 'jpskgc-article',
  storageBucket: '',
  messagingSenderId: '435581300302',
  appId: '1:435581300302:web:698faafeed154e99',
};

firebase.initializeApp(firebaseConfig);

export default firebase;
