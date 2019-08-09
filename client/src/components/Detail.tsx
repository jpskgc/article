import React from 'react';
import {RouteComponentProps, withRouter} from 'react-router';
import {Container, Header} from 'semantic-ui-react';
import axios from 'axios';
import {Article} from '../articleData';

interface ArticleState {
  article: Article;
}

class Detail extends React.Component<
  RouteComponentProps<{id: string}>,
  ArticleState
> {
  constructor(props: RouteComponentProps<{id: string}>) {
    super(props);
    this.state = {
      article: {
        id: 0,
        title: '',
        content: '',
        imageNames: [],
      },
    };
    this.serverRequest = this.serverRequest.bind(this);
  }

  serverRequest() {
    // const client = axios.create({
    //   baseURL: process.env.REACT_APP_API_URL,
    // });
    // client
    axios
      .get('/api/article/' + this.props.match.params.id)
      .then(response => {
        this.setState({article: response.data});
      })
      .catch(response => console.log('ERROR!! occurred in Backend.'));
  }

  componentDidMount() {
    this.serverRequest();
  }

  Paragraph = () => <p>{[this.state.article.content].join('')}</p>;

  render() {
    //TODO redirect
    return (
      <Container text style={{marginTop: '3em'}}>
        <Header as="h1">{this.state.article.title}</Header>
        <this.Paragraph />
        {(this.state.article.imageNames || []).map(function(articleData, i) {
          return (
            <img
              src={`https://article-s3-jpskgc.s3-ap-northeast-1.amazonaws.com/media/${
                articleData.name
              }`}
              alt="new"
            />
          );
        })}
      </Container>
    );
  }
}

export default withRouter(Detail);
