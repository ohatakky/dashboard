import React, { FC } from "react";
import Link from "@material-ui/core/Link";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Title from "~/components/common/Title";

const useStyles = makeStyles({
  cardContext: {
    flex: 1,
  },
});

const Card: FC<{ title: string }> = ({ title }) => {
  const classes = useStyles();
  return (
    <React.Fragment>
      <Title>{title}</Title>
      <Typography component="p" variant="h4">
        {}
      </Typography>
      <Typography color="textSecondary" className={classes.cardContext}>
        {}
      </Typography>
      <div>
        <Link color="primary" href="#" onClick={(e) => e.preventDefault()}>
          {}
        </Link>
      </div>
    </React.Fragment>
  );
};

export default Card;
