import React, { FC } from "react";
import {
  Router as ReactRouter,
  Switch,
  Route,
  Redirect,
} from "react-router-dom";
import { History } from "history";
import Layout from "~/components/common/Layout";
import Summary from "~/components/pages/Summary";
import Analytics from "~/components/pages/Analytics";

const Router: FC<{ history: History }> = ({ history }) => (
  <ReactRouter history={history}>
    <Switch>
      <Route path="/summary" exact>
        <Layout title="Summary">
          <Summary />
        </Layout>
      </Route>
      <Route path="/analytics" exact>
        <Layout title="Analytics">
          <Analytics />
        </Layout>
      </Route>
      <Route path="/">
        <Redirect
          to={{
            pathname: "/summary",
          }}
        />
      </Route>
    </Switch>
  </ReactRouter>
);

export default Router;
