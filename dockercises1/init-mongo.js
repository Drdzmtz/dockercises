db = db.getSiblingDB('dockercises1')
db.auth('tredicom', 'tredicom')

db.createUser(
    {
        user: 'tredicom-user',
        pwd: 'tredicom',
        roles: [
            {
                role: 'readWrite',
                db: 'dockercises1'
            }
        ]
    }
)

db.createCollection('Personas')