import React from 'react';
import {Container, Message, Button, Icon} from 'semantic-ui-react';

class Finish extends React.Component<{}, {}> {
  render() {
    return (
      <Container text style={{marginTop: '3em'}}>
        <Message color="green">Your post was successful</Message>
        <Button color="green" as="a" href="/">
          <Icon name="arrow left" />
          Back to Home
        </Button>
      </Container>
    );
  }
}

export default Finish;
