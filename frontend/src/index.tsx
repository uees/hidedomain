import React from 'react';
import ReactDOM from 'react-dom/client';
import { RouterProvider } from "react-router-dom";
import { observer } from 'mobx-react-lite';

import reportWebVitals from './reportWebVitals';
import router from './router';


const App: React.FC = observer(() => {
  return (
    <React.StrictMode >
      <RouterProvider router={router} />
    </React.StrictMode >
  )
});

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

root.render(
  <App />
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
