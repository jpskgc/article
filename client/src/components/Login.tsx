import React from 'react';
import {Button, Form, Grid, Header, Message, Segment} from 'semantic-ui-react';
import firebase from 'firebase';
import {Redirect} from 'react-router';

interface userState {
  email: string;
  password: string;
  redirect: boolean;
}

class LoginForm extends React.Component<{}, userState> {
  constructor(props: {}) {
    super(props);
    this.state = {
      email: '',
      password: '',
      redirect: false,
    };
    this.login = this.login.bind(this);
    this.handleChangeEmail = this.handleChangeEmail.bind(this);
    this.handleChangePassword = this.handleChangePassword.bind(this);
    this.renderRedirect = this.renderRedirect.bind(this);
  }

  handleChangeEmail(e: React.FormEvent<HTMLInputElement>) {
    this.setState({email: e.currentTarget.value});
  }

  handleChangePassword(e: React.FormEvent<HTMLInputElement>) {
    this.setState({password: e.currentTarget.value});
  }

  renderRedirect = () => {
    if (this.state.redirect) {
      return <Redirect to="" />;
    }
  };

  login() {
    firebase
      .auth()
      .signInWithEmailAndPassword(this.state.email, this.state.password)
      .catch(error => {
        console.log(error.message);
      });
    this.setState({
      redirect: true,
    });
  }
  render() {
    return (
      <Grid textAlign="center" style={{height: '100vh'}} verticalAlign="middle">
        <Grid.Column style={{maxWidth: 450}}>
          <Header as="h2" color="teal" textAlign="center">
            Log-in to your account
          </Header>
          <Form size="large" onSubmit={this.login}>
            <Segment stacked>
              <Form.Input
                fluid
                icon="user"
                iconPosition="left"
                placeholder="E-mail address"
                value={this.state.email}
                onChange={this.handleChangeEmail}
              />
              <Form.Input
                fluid
                icon="lock"
                iconPosition="left"
                placeholder="Password"
                type="password"
                value={this.state.password}
                onChange={this.handleChangePassword}
              />
              {this.renderRedirect()}
              <Button color="teal" fluid size="large">
                Login
              </Button>
            </Segment>
          </Form>
          <Message>
            New to us? <a href="/signup">Sign Up</a>
          </Message>
          <Message>
            <Message.Header>Test User Login</Message.Header>
            <Message.List>
              <Message.Item>test@test.co.jp</Message.Item>
              <Message.Item>test123</Message.Item>
            </Message.List>
          </Message>
        </Grid.Column>
      </Grid>
    );
  }
}

export default LoginForm;
