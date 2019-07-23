import React from 'react';
import {Redirect, Route, Switch} from 'react-router';
import FixedMenuLayout from './components/Menu';
import List from './components/List';
import Detail from './components/Detail';
import Post from './components/Post';
import Finish from './components/Finish';
import Login from './components/Login';
import Signup from './components/Signup';
import PrivateRouter from './components/PrivateRouter';
import Setting from './components/Setting';

class App extends React.Component {
  render() {
    return (
      <div className="container">
        <FixedMenuLayout />
        <Switch>
          <Route exact path="/detail/:id" component={Detail} />
          <Route path="/post/finish" component={Finish} />
          <Route path="/post" component={Post} />
          <Route path="/login" component={Login} />
          <Route path="/signup" component={Signup} />
          <PrivateRouter path="/setting" component={Setting} />
          <Route path="/" component={List} />
          <Redirect to="/" />;
        </Switch>
      </div>
    );
  }
}
export default App;
