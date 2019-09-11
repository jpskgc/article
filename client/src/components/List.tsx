import React from 'react';
import {RouteComponentProps, withRouter, Redirect} from 'react-router';
import {Link} from 'react-router-dom';
import {
  Grid,
  Segment,
  Container,
  Header,
  Pagination,
  PaginationProps,
} from 'semantic-ui-react';
import {Article} from '../articleData';
import {getArticleFactory} from '../api/articleAPI';

interface ArticleState {
  articleFromApi: Article[];
  articleToDisplay: Article[];
  begin: number;
  end: number;
  activePage: number;
  redirect: boolean;
}

class List extends React.Component<
  RouteComponentProps<{id: string}>,
  ArticleState
> {
  constructor(props: RouteComponentProps<{id: string}>) {
    super(props);
    this.state = {
      articleFromApi: [],
      articleToDisplay: [],
      begin: 0,
      end: 5,
      activePage: 1,
      redirect: false,
    };
    this.getArticle = this.getArticle.bind(this);
    this.pageChange = this.pageChange.bind(this);
    this.setActiveList = this.setActiveList.bind(this);
  }

  async getArticle() {
    const getArticle = getArticleFactory();
    const articleData = await getArticle();

    if (articleData == null) {
      this.setState({articleFromApi: []});
    } else {
      this.setState({articleFromApi: articleData});
    }
  }

  async pageChange(
    event: React.MouseEvent<HTMLAnchorElement>,
    data: PaginationProps
  ) {
    await this.setState({activePage: data.activePage as number});
    await this.setActiveList();
    this.setState({redirect: true});
  }

  async setActiveList() {
    await this.setState({begin: this.state.activePage * 5 - 5});
    await this.setState({end: this.state.activePage * 5});
    this.setState(
      {
        articleToDisplay: this.state.articleFromApi.slice(
          this.state.begin,
          this.state.end
        ),
      },
      () => window.scrollTo(0, 0)
    );
  }

  renderRedirect = () => {
    if (this.state.redirect) {
      return <Redirect to={`/${this.state.activePage}`} />;
    }
  };

  async componentDidMount() {
    this.setState({articleFromApi: []});
    await this.getArticle();
    await this.setState({activePage: Number(this.props.match.params.id)});
    this.setActiveList();
  }

  render() {
    return (
      <Container style={{marginTop: '3em'}} text>
        <Grid columns={1} divided="vertically">
          <Grid.Row>
            {(this.state.articleToDisplay || []).map(function(articleData, i) {
              return (
                <Grid.Column>
                  <Segment>
                    <Header as="h1">{articleData.title}</Header>
                    <p style={{textOverflow: 'clip', wordBreak: 'break-all'}}>
                      {articleData.content.length > 100
                        ? articleData.content.substring(0, 97) + '...'
                        : articleData.content}
                    </p>
                    <Link to={`/detail/${articleData.id}`}>
                      continue reading
                    </Link>
                  </Segment>
                </Grid.Column>
              );
            })}
          </Grid.Row>
        </Grid>
        <Pagination
          defaultActivePage={this.props.match.params.id}
          totalPages={Math.ceil(this.state.articleFromApi.length / 5)}
          onPageChange={this.pageChange}
        />
        {this.renderRedirect()}
      </Container>
    );
  }
}

export default withRouter(List);
