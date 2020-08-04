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

const Card: FC<{ title: string; average: number; correlation?: number }> = ({
  title,
  average,
  correlation,
}) => {
  const classes = useStyles();
  return (
    <React.Fragment>
      <Title>{title}</Title>
      <Typography component="p" variant="h5">
        Average
      </Typography>
      <Typography color="textSecondary" className={classes.cardContext}>
        {`${average.toFixed(2)} h`}
      </Typography>
      <Typography component="p" variant="h5">
        Correlation Coefficient
      </Typography>
      <Typography color="textSecondary" className={classes.cardContext}>
        {correlation ? correlation : ""}
      </Typography>
    </React.Fragment>
  );
};

export default Card;
