import React from 'react';
import {BrowserRouter as Router} from 'react-router-dom';
import Menu from './Menu';
import renderer from 'react-test-renderer';

it('renders correctly', () => {
  const tree = renderer
    .create(
      <Router>
        <Menu />
      </Router>
    )
    .toJSON();
  expect(tree).toMatchSnapshot();
});
