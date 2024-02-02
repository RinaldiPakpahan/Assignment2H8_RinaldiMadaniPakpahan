var input = ['a', 'b', 'a', 1, 1, 2, 1, '1'];

function cek(arr) {
    var result = arr.map(function (element, index) {
        var count = arr.filter(function (el) {
            return el === element;
        }).length;

        return { element: element, count: count };
    });

    var maxCount = Math.max(...result.map(item => item.count));
    var mostFrequent = result.find(item => item.count === maxCount);

    console.log(mostFrequent.element, mostFrequent.count);
}

cek(input);