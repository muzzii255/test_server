import random
import time

import requests

BASE_URL = "http://localhost:9000"


# Sample data generators
def random_user():
    return {
        "name": random.choice(["Alice", "Bob", "Charlie", "Diana", "Eve"]),
        "age": random.randint(18, 65),
        "contactInfo": {
            "email": f"user{random.randint(1000, 9999)}@example.com",
            "phone": f"+1875{random.randint(1000245, 9999999)}",
        },
        "address": [
            {"street": "your mom blvd", "zipcode": f"{random.randint(10001, 99999)}"},
            {"street": "your mom blvd", "zipcode": f"{random.randint(10001, 99999)}"},
            {"street": "your mom blvd", "zipcode": f"{random.randint(10001, 99999)}"},
            {"street": "your mom blvd", "zipcode": f"{random.randint(10001, 99999)}"},
            {"street": "your mom blvd", "zipcode": f"{random.randint(10001, 99999)}"},
        ],
    }


def random_product():
    products = ["Laptop", "Phone", "Tablet", "Monitor", "Keyboard", "Mouse"]
    return {
        "name": random.choice(products),
        "description": f"High quality {random.choice(products).lower()}",
        "price": round(random.uniform(10.0, 1000.0), 2),
        "stock": random.randint(0, 100),
    }


endpoints = [
    {
        "name": "users",
        "base": "/api/v1/users",
        "generator": random_user,
        "methods": ["POST", "GET", "PUT", "DELETE"],
    }
]

created_ids = {
    "users": [],
    "products": [],
}


def make_request(endpoint, method):
    url = BASE_URL + endpoint["base"]

    try:
        if method == "POST":
            data = endpoint["generator"]()
            resp = requests.post(url, json=data)
            print(f"POST {url} - {resp.status_code}")
            if resp.status_code in [200, 201]:
                try:
                    json_data = resp.json()
                    if "id" in json_data:
                        created_ids[endpoint["name"]].append(json_data["id"])
                except:
                    pass

        elif method == "GET":
            if random.choice([True, False]) and created_ids[endpoint["name"]]:
                # Get specific ID
                id_val = random.choice(created_ids[endpoint["name"]])
                url = f"{url}/{id_val}"
                resp = requests.get(url)
                print(f"GET {url} - {resp.status_code}")
            else:
                resp = requests.get(url)
                print(f"GET {url} - {resp.status_code}")

        elif method == "PUT":
            if created_ids[endpoint["name"]]:
                id_val = random.choice(created_ids[endpoint["name"]])
                url = f"{url}/{id_val}"
                data = endpoint["generator"]()
                resp = requests.put(url, json=data)
                print(f"PUT {url} - {resp.status_code}")
            else:
                print(f"PUT {url} - Skipped (no IDs)")

        elif method == "DELETE":
            if created_ids[endpoint["name"]]:
                id_val = random.choice(created_ids[endpoint["name"]])
                url = f"{url}/{id_val}"
                resp = requests.delete(url)
                print(f"DELETE {url} - {resp.status_code}")
                created_ids[endpoint["name"]].remove(id_val)
            else:
                print(f"DELETE {url} - Skipped (no IDs)")

    except Exception as e:
        print(f"Error on {method} {url}: {e}")


def main():
    print("Starting random requests to endpoints...")
    print(f"Target: {BASE_URL}\n")

    print("=== Creating initial data ===")
    for endpoint in endpoints:
        for _ in range(4):
            make_request(endpoint, "POST")
            time.sleep(0.1)

    print("\n=== Random requests ===")
    for _ in range(6):
        endpoint = random.choice(endpoints)
        method = random.choice(endpoint["methods"])
        make_request(endpoint, method)
        time.sleep(random.uniform(0.1, 0.5))

    print("\n=== Done ===")
    print(f"Created IDs: {created_ids}")


if __name__ == "__main__":
    main()
