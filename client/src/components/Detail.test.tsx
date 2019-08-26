import React from 'react';
import {BrowserRouter as Router} from 'react-router-dom';
import Detail from './Detail';
import renderer from 'react-test-renderer';

it('renders correctly', () => {
  const tree = renderer
    .create(
      <Router>
        <Detail />
      </Router>
    )
    .toJSON();
  expect(tree).toMatchSnapshot();
});
