using MySql.Data;
using MySql.Data.MySqlClient;

namespace Metel_Template.code
{
    public class DBManager
    {
        public MySqlConnection connection;
        private string server;
        private string database;
        private string uid;
        private string password;

        public DBManager()
        {
            Initialize();
        }

        private void Initialize()
        {
            server = "ipadress";
            database = "database";
            uid = "username";
            password = "password";
            string connectionString;
            connectionString = $"SERVER={server};DATABASE={database};UID={uid};PASSWORD={password};";

            connection = new MySqlConnection(connectionString);
        }

        public bool OpenConnection()
        {
            try
            {
                connection.Open();
                return true;
            }
            catch (MySqlException ex)
            {
                switch (ex.Number)
                {
                    case 0:
                        MessageBox.Show("Error! Could not connect to the server. Contact an admin!");
                        break;
                    case 1:
                        MessageBox.Show("Error! Username/password invalid!");
                        break;
                }
                return false;
            }
        }

        public bool CloseConnection()
        {
            try
            {
                connection.Close(); 
                return true;
            }
            catch (MySqlException ex)
            {
                MessageBox.Show(ex.Message); 
                return false;
            }
        }

        public void Update(string query)
        {
            if (this.OpenConnection() == true)
            {
                MySqlCommand cmd = new MySqlCommand(query, connection);
                cmd.ExecuteNonQuery();
                this.CloseConnection();
            }
        }
    }
}