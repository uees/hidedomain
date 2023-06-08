import React from 'react';
import { storeContext } from '../store';

const useStore = () => React.useContext(storeContext);

export default useStore;
