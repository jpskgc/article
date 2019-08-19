import React from 'react';
import {RouteComponentProps, withRouter} from 'react-router';
import {Container, Header, Button, Icon} from 'semantic-ui-react';
import axios from 'axios';
import {Article} from '../articleData';
import {Redirect} from 'react-router';
import {Link} from 'react-router-dom';
import firebase from '../firebase/firebase';

interface ArticleState {
  article: Article;
  redirect: boolean;
  user: firebase.User | null;
}

// interface UserStatus {
//   user: firebase.User | null;
// }

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
      user: null,
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
    this.setState(() => window.scrollTo(0, 0));
    firebase.auth().onAuthStateChanged(user => {
      this.setState({user});
    });
  }

  Paragraph = () => (
    <p style={{whiteSpace: 'pre-wrap', wordBreak: 'break-all'}}>
      {[this.state.article.content].join('')}
    </p>
  );

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
        {this.renderRedirect()}
        <Container tyle={{display: 'flex'}}>
          <Link to="/">
            <Button color="green">
              <Icon name="arrow left" />
              Home
            </Button>
          </Link>
          {this.state.user ? (
            <Button floated="right" onClick={this.deleteArticle}>
              <Icon name="trash" />
              Delete this Article
            </Button>
          ) : (
            <div />
          )}
        </Container>
      </Container>
    );
  }
}

export default withRouter(Detail);
