import React from 'react';
import {Button, Form, Grid, Header, Message, Segment} from 'semantic-ui-react';
import firebase from '../firebase/firebase';
import {Redirect} from 'react-router';

interface userState {
  email: string;
  password: string;
  redirect: boolean;
}

class SignupForm extends React.Component<{}, userState> {
  constructor(props: {}) {
    super(props);
    this.state = {
      email: '',
      password: '',
      redirect: false,
    };
    this.signUp = this.signUp.bind(this);
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

  signUp() {
    firebase
      .auth()
      .createUserWithEmailAndPassword(this.state.email, this.state.password)
      .then(res => {
        console.log('Create account: ', res.user);
      })
      .catch(error => {
        console.log(error.message);
      });
    this.setState({
      redirect: true,
    });
  }

  renderRedirect = () => {
    if (this.state.redirect) {
      return <Redirect to="" />;
    }
  };

  render() {
    return (
      <Grid textAlign="center" style={{height: '100vh'}} verticalAlign="middle">
        <Grid.Column style={{maxWidth: 450}}>
          <Header as="h2" color="teal" textAlign="center">
            Sign-up to your account
          </Header>
          <Form size="large" onSubmit={this.signUp}>
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
                Signup
              </Button>
            </Segment>
          </Form>
          <Message>
            already have account? <a href="/login">Login</a>
          </Message>
        </Grid.Column>
      </Grid>
    );
  }
}

export default SignupForm;
