import React from 'react';
import Finish from './Finish';
import renderer from 'react-test-renderer';

it('renders correctly', () => {
  const tree = renderer.create(<Finish />).toJSON();
  expect(tree).toMatchSnapshot();
});
