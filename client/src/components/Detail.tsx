import React from 'react';
import {RouteComponentProps, withRouter} from 'react-router';
import {Container, Header, Button, Icon} from 'semantic-ui-react';
import {Article} from '../articleData';
import {Redirect} from 'react-router';
import firebase from '../firebase/firebase';
import {getSingleArticleFactory, deleteArticleFactory} from '../api/articleAPI';

interface ArticleState {
  article: Article;
  redirect: boolean;
  user: firebase.User | null;
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
      user: null,
    };
    this.getArticle = this.getArticle.bind(this);
    this.deleteArticle = this.deleteArticle.bind(this);
  }

  async getArticle() {
    const getSingleArticle = getSingleArticleFactory();
    const articleData = await getSingleArticle(this.props.match.params.id);
    this.setState({
      article: articleData,
    });
  }

  async deleteArticle() {
    this.setState({redirect: true});
    const deleteArticle = deleteArticleFactory();
    const response = await deleteArticle(this.props.match.params.id);
    this.setState({article: response});
  }

  renderRedirect = () => {
    if (this.state.redirect) {
      return <Redirect to="/article/delete/finish" />;
    }
  };

  componentDidMount() {
    this.getArticle();
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
          <Button color="green" onClick={() => this.props.history.goBack()}>
            <Icon name="arrow left" />
            Home
          </Button>
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
