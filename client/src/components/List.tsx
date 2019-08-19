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
import axios from 'axios';
import {Article} from '../articleData';

interface ArticleState {
  articles: Article[];
  articleDatas: Article[];
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
      articles: [],
      articleDatas: [],
      begin: 0,
      end: 5,
      activePage: 1,
      redirect: false,
    };
    this.serverRequest = this.serverRequest.bind(this);
    this.btnClick = this.btnClick.bind(this);
    this.setList = this.setList.bind(this);
  }

  async serverRequest() {
    const res = await axios.get('/api/articles');
    if (res.data == null) {
      this.setState({articles: []});
    } else {
      this.setState({articles: res.data});
    }
  }

  async btnClick(
    event: React.MouseEvent<HTMLAnchorElement>,
    data: PaginationProps
  ) {
    await this.setState({activePage: data.activePage as number});
    await this.setState({begin: this.state.activePage * 5 - 5});
    await this.setState({end: this.state.activePage * 5});
    this.setState(
      {
        articleDatas: this.state.articles.slice(
          this.state.begin,
          this.state.end
        ),
      },
      () => window.scrollTo(0, 0)
    );
    this.setState({redirect: true});
  }

  async setList() {
    await this.setState({activePage: Number(this.props.match.params.id)});
    await this.setState({begin: this.state.activePage * 5 - 5});
    await this.setState({end: this.state.activePage * 5});
    this.setState(
      {
        articleDatas: this.state.articles.slice(
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
    this.setState({articles: []});
    await this.serverRequest();
    await this.setState({
      articleDatas: this.state.articles.slice(this.state.begin, this.state.end),
    });
    this.setList();
  }

  render() {
    return (
      <Container style={{marginTop: '3em'}} text>
        <Grid columns={1} divided="vertically">
          <Grid.Row>
            {(this.state.articleDatas || []).map(function(articleData, i) {
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
          totalPages={Math.ceil(this.state.articles.length / 5)}
          onPageChange={this.btnClick}
        />
        {this.renderRedirect()}
      </Container>
    );
  }
}

export default withRouter(List);
