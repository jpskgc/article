import React from 'react';
import {Container, Message} from 'semantic-ui-react';

class Finish extends React.Component<{}, {}> {
  render() {
    return (
      <Container text style={{marginTop: '3em'}}>
        <Message color="green">Your post was successful</Message>
      </Container>
    );
  }
}

export default Finish;
