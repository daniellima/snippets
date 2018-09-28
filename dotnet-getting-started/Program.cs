using System;
using Newtonsoft.Json;

namespace dotnet_getting_started
{
    public class Account {
        public string Name { get; set; }
        public string Email { get; set; }
        public DateTime DOB { get; set; }
    }

    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");

            Account account = new Account {
                Name = "John Doe",
                Email = "john@example.com",
                DOB = new DateTime(1990, 2, 20, 0, 0, 0, DateTimeKind.Utc),
            };

            string json = JsonConvert.SerializeObject(account, Formatting.Indented);
            Console.WriteLine(json);
        }
    }
}
