import React from 'react';
import {Container, Menu} from 'semantic-ui-react';
import firebase from '../firebase/firebase';
import {Link} from 'react-router-dom';

interface UserStatus {
  user: firebase.User | null;
}

class FixedMenuLayout extends React.Component<{}, {}> {
  public state: UserStatus = {
    user: null,
  };

  componentDidMount() {
    firebase.auth().onAuthStateChanged(user => {
      this.setState({user});
    });
  }

  logout() {
    firebase.auth().signOut();
  }

  render() {
    return (
      <div>
        <Menu fixed="top" inverted>
          <Container>
            <Menu.Item as="a" header>
              Project Name
            </Menu.Item>
            <Menu.Item as="a" href="/">
              Home
            </Menu.Item>
            {this.state.user ? (
              // <Menu.Item as="a" href="/post">
              //   Post
              // </Menu.Item>
              <Menu.Item name="post">
                <Link to="/post">Post</Link>
              </Menu.Item>
            ) : (
              <div />
            )}
            {this.state.user ? (
              <Menu.Menu position="right">
                <Menu.Item name="setting">
                  {/* TODO 選択範囲 */}
                  <Link to="/setting">Setting</Link>
                </Menu.Item>
                <Menu.Item name="logout" onClick={this.logout}>
                  <Link to="">Log-out</Link>
                </Menu.Item>
              </Menu.Menu>
            ) : (
              <Menu.Menu position="right">
                <Menu.Item name="login" href="/login">
                  Log-in
                </Menu.Item>
              </Menu.Menu>
            )}
          </Container>
        </Menu>
      </div>
    );
  }
}

export default FixedMenuLayout;
