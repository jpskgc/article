import React from 'react';
import {BrowserRouter as Router} from 'react-router-dom';
import List from './List';
import renderer from 'react-test-renderer';

it('renders correctly', () => {
  const tree = renderer
    .create(
      <Router>
        <List />
      </Router>
    )
    .toJSON();
  expect(tree).toMatchSnapshot();
});
