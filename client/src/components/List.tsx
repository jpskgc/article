import React from 'react';
import {Link} from 'react-router-dom';
import {Grid, Segment, Container, Header} from 'semantic-ui-react';
import axios from 'axios';
import {Article} from '../articleData';

interface ArticleState {
  articles: Article[];
}

class List extends React.Component<{}, ArticleState> {
  constructor(props: {}) {
    super(props);
    this.state = {
      articles: [],
    };
    this.serverRequest = this.serverRequest.bind(this);
  }

  serverRequest() {
    // const client = axios.create({
    //   baseURL: process.env.REACT_APP_API_URL,
    // });
    // client
    axios
      .get('/api/articles')
      .then(response => {
        this.setState({articles: response.data});
      })
      .catch(response => console.log('ERROR!! occurred in Backend.'));
  }

  componentDidMount() {
    this.setState({articles: []});
    this.serverRequest();
  }

  render() {
    return (
      <Container style={{marginTop: '7em'}} text>
        <Grid columns={1} divided="vertically">
          <Grid.Row>
            {(this.state.articles || []).map(function(articleData, i) {
              return (
                <Grid.Column>
                  <Segment>
                    <Header as="h1">{articleData.title}</Header>
                    <p>{articleData.content}</p>
                    <Link to={`/detail/${articleData.id}`}>
                      continue reading
                    </Link>
                  </Segment>
                </Grid.Column>
              );
            })}
          </Grid.Row>
        </Grid>
      </Container>
    );
  }
}

export default List;
