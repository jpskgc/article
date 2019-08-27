import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';
import * as api from './articleAPI';
import articlesMockData from './__mocks__/articles.json';
import singleArticleMockData from './__mocks__/singleArticle.json';

describe('article API handlers', () => {
  const mock = new MockAdapter(axios);
  afterEach(() => {
    mock.reset();
  });
  describe('get articles', () => {
    it('should succeed', async () => {
      mock.onGet(`/api/articles`).reply(200, articlesMockData);
      const getArticle = api.getArticleFactory();
      const articles = await getArticle();
      expect(articles[0].title).toBe(articlesMockData[0].title);
    });
  });
  describe('get single article', () => {
    it('should succeed', async () => {
      mock.onGet('/api/article/' + 1).reply(200, singleArticleMockData);
      const getSingleArticle = api.getSingleArticleFactory();
      const article = await getSingleArticle('1');
      expect(article.title).toBe(singleArticleMockData.title);
    });
  });
  describe('delete single article', () => {
    it('should succeed', async () => {
      mock.onGet('/api/delete/' + 1).reply(200, {status: 'ok'});
      const deleteArticle = api.deleteArticleFactory();
      const article = await deleteArticle('1');
      expect(article.status).toBe({status: 'ok'}.status);
    });
  });
  describe('post article', () => {
    it('should succeed', async () => {
      mock
        .onPost('/api/post', {
          title: 'aaa',
          content: 'aaa',
        })
        .reply(200, {uuid: 'bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c'});
      const postArticle = api.postArticleFactory();
      const article = await postArticle({
        title: 'aaa',
        content: 'aaa',
      });
      expect(article.uuid).toBe('bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c');
    });
  });
  describe('post article image', () => {
    it('should succeed', async () => {
      mock
        .onPost('/api/post/image', new FormData())
        .reply(200, [{name: '1925c2de071aff40eca2ac15524fe139-300x300.jpg'}]);
      const postArticleImage = api.postArticleImageFactory();
      const article = await postArticleImage(new FormData());
      expect(article[0].name).toBe(
        '1925c2de071aff40eca2ac15524fe139-300x300.jpg'
      );
    });
  });
  describe('post article image data to DB', () => {
    it('should succeed', async () => {
      mock
        .onPost('/api/post/image/db', {
          articleUUID: 'bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c',
          imageNames: 'a70f2ef9-301f-4ed1-n4b4-f61d64e9f82a',
        })
        .reply(200);
      const postArticleImageToDB = api.postArticleImageToDBFactory();
      const article = await postArticleImageToDB({
        articleUUID: 'bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c',
        imageNames: 'a70f2ef9-301f-4ed1-n4b4-f61d64e9f82a',
      });
      expect(article.status).toBe(200);
    });
  });
});
