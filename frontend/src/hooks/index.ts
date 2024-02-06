import React from "react";
import { storeContext } from "../store";

export const useStore = () => React.useContext(storeContext);

export const useTitle = (title: string) => {
  const { site } = useStore();

  React.useEffect(() => {
    const prevTitle = document.title;
    document.title = title;

    return () => {
      document.title = prevTitle;
    };
  }, [site, title]);
};
