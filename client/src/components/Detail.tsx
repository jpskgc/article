import React from 'react';
import {RouteComponentProps, withRouter} from 'react-router';
import {Container, Header, Button, Icon} from 'semantic-ui-react';
import axios from 'axios';
import {Article} from '../articleData';
import {Redirect} from 'react-router';

interface ArticleState {
  article: Article;
  redirect: boolean;
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
      redirect: false,
    };
    this.serverRequest = this.serverRequest.bind(this);
    this.deleteArticle = this.deleteArticle.bind(this);
  }

  serverRequest() {
    axios
      .get('/api/article/' + this.props.match.params.id)
      .then(response => {
        this.setState({article: response.data});
      })
      .catch(response => console.log('ERROR!! occurred in Backend.'));
  }

  deleteArticle() {
    this.setState({redirect: true});
    axios
      .get('/api/delete/' + this.props.match.params.id)
      .then(response => {
        this.setState({article: response.data});
      })
      .catch(response => console.log('ERROR!! occurred in Backend.'));
  }

  renderRedirect = () => {
    if (this.state.redirect) {
      return <Redirect to="/article/delete/finish" />;
    }
  };

  componentDidMount() {
    this.serverRequest();
  }

  Paragraph = () => <p>{[this.state.article.content].join('')}</p>;

  render() {
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
        {/* TODO fix button location when image exists */}
        <Button color="green" as="a" href="/">
          <Icon name="arrow left" />
          Back to Home
        </Button>
        {this.renderRedirect()}
        <Button floated="right" onClick={this.deleteArticle}>
          <Icon name="trash" />
          Delete this Article
        </Button>
      </Container>
    );
  }
}

export default withRouter(Detail);
