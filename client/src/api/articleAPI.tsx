import axios from 'axios';
import {Article} from '../articleData';

export const getArticleFactory = () => {
  const getArticle = async () => {
    try {
      const response = await axios.get('/api/articles');

      if (response.status !== 200) {
        throw new Error('Server Error');
      }
      const article: Article[] = response.data;

      return article;
    } catch (err) {
      throw err;
    }
  };
  return getArticle;
};

export const getSingleArticleFactory = () => {
  const getSingleArticle = async (id: string) => {
    try {
      const response = await axios.get('/api/article/' + id);

      if (response.status !== 200) {
        throw new Error('Server Error');
      }
      const article: Article = response.data;
      return article;
    } catch (err) {
      throw err;
    }
  };
  return getSingleArticle;
};

export const deleteArticleFactory = () => {
  const deleteArticle = async (id: string) => {
    try {
      const response = await axios.get('/api/delete/' + id);

      if (response.status !== 200) {
        throw new Error('Server Error');
      }

      return response.data;
    } catch (err) {
      throw err;
    }
  };
  return deleteArticle;
};

export const postArticleFactory = () => {
  const postArticle = async (articleTextData: {
    title: string;
    content: string;
  }) => {
    try {
      const response = await axios.post('/api/post', articleTextData);

      if (response.status !== 200) {
        throw new Error('Server Error');
      }

      return response.data;
    } catch (err) {
      throw err;
    }
  };
  return postArticle;
};

export const postArticleImageFactory = () => {
  const postArticleImage = async (formData: FormData) => {
    try {
      const response = await axios.post('/api/post/image', formData, {
        headers: {'Content-Type': 'multipart/form-data'},
      });

      if (response.status !== 200) {
        throw new Error('Server Error');
      }

      return response.data;
    } catch (err) {
      throw err;
    }
  };
  return postArticleImage;
};

export const postArticleImageToDBFactory = () => {
  const postArticleImageToDB = async (imageData: {
    articleUUID: any;
    imageNames: any;
  }) => {
    try {
      const response = await axios.post('/api/post/image/db', imageData);

      if (response.status !== 200) {
        throw new Error('Server Error');
      }

      return response;
    } catch (err) {
      throw err;
    }
  };
  return postArticleImageToDB;
};
