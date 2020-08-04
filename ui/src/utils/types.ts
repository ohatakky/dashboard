export type Time = {
  time: string;
  valid: boolean;
};

export type String = {
  string: string;
  valid: boolean;
};

export type Int = {
  int: number;
  valid: boolean;
};

export type Float = {
  float: number;
  valid: boolean;
};

export type Bool = {
  bool: boolean;
  valid: boolean;
};

export type Life = {
  date: Time;
  condition: Int;
  rising: Time;
  sleep: Float;
  light_off: Bool;
  bath: Time;
  fullness: Int;
  vitamin: Bool;
  weather: String;
  hunting: Float;
  devotion: Float;
  hobby: Float;
  workout_w: Float;
  workout_r: Float;
  workout_b: Int;
};
