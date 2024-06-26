import { buildQuery } from "../../../services/graphql-services";
function useGraphql() {
  const queryKey = "Movie";
  const getAllJobTitles = buildQuery({
    operation: "Movies",
    params:{input:"ListMovieInput!"},
    options: {
      type: "query",
    },
    node: `
        data
        {
          id
          title
          genre
          status
          language
          director
          cast
          poster
          rated
          duration
          trailer
          openingDay
          story
        }
        pagination{
          total
        }
        `,
  });


  const createJobTitle = buildQuery({
    operation: "CreateJobTitle",
    options: {
      type: "mutation",
    },
    node: `
          data {
            id
          }
        `,
    params: {
      input: "NewJobTitleInput!",
    },
  });

  return { getAllJobTitles, createJobTitle, queryKey };
}
export default useGraphql;
