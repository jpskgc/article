import React from 'react';
import {Route, Redirect} from 'react-router-dom';
import firebase from '../firebase/firebase';

function PrivateRoute({component: Component, ...rest}: any) {
  return (
    <Route
      {...rest}
      render={props =>
        firebase.auth().currentUser !== null ? (
          <Component {...props} />
        ) : (
          <Redirect to="" />
        )
      }
    />
  );
}

export default PrivateRoute;
