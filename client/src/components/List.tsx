import React from 'react';
import {Link} from 'react-router-dom';
import {
  Grid,
  Segment,
  Container,
  Header,
  Pagination,
  PaginationProps,
  Icon,
} from 'semantic-ui-react';
import axios from 'axios';
import {Article} from '../articleData';

interface ArticleState {
  articles: Article[];
  articleDatas: Article[];
  begin: number;
  end: number;
  activePage: number;
}

class List extends React.Component<{}, ArticleState> {
  constructor(props: {}) {
    super(props);
    this.state = {
      articles: [],
      articleDatas: [],
      begin: 0,
      end: 5,
      activePage: 1,
    };
    this.serverRequest = this.serverRequest.bind(this);
    this.btnClick = this.btnClick.bind(this);
  }

  async serverRequest() {
    const res = await axios.get('/api/articles');
    this.setState({articles: res.data});
  }

  async btnClick(
    event: React.MouseEvent<HTMLAnchorElement>,
    data: PaginationProps
  ) {
    await this.setState({activePage: data.activePage as number});
    await this.setState({begin: this.state.activePage * 5 - 5});
    await this.setState({end: this.state.activePage * 5});
    this.setState({
      articleDatas: this.state.articles.slice(this.state.begin, this.state.end),
    });
  }

  async componentDidMount() {
    this.setState({articles: []});
    await this.serverRequest();
    this.setState({
      articleDatas: this.state.articles.slice(this.state.begin, this.state.end),
    });
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
        <Pagination
          defaultActivePage={1}
          totalPages={Math.ceil(this.state.articles.length / 5)}
          onPageChange={this.btnClick}
        />
      </Container>
    );
  }
}

export default List;
