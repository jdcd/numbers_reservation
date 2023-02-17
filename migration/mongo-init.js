use("numbers_reservation")
db.reservations.drop()

db.createCollection('reservations',{
    validator: {
        $jsonSchema:{
            bsonType: 'object',
            required: ['reservation_number'],
            properties:{
                reservation_number: {
                    bsonType:'number'
                }
            }
        }
    }
})

db.reservations.createIndex({reservation_number: 1}, { unique: true })
