export type postType = {
  title: String;
  user: {
    name: String;
    major: {
      name: String;
      faculty: {
        name: String;
      };
    };
    imageUrl: String;
  };
  tags: String[];
  createdAt: Date;
};

export type userType = {
  name: String;
  postCount: Number;
  bookmarkCount: Number;
};
