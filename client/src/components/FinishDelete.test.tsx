import React from 'react';
import FinishDetele from './FinishDelete';
import renderer from 'react-test-renderer';

it('renders correctly', () => {
  const tree = renderer.create(<FinishDetele />).toJSON();
  expect(tree).toMatchSnapshot();
});
