export const capitalizeFirstLetter = (string: string) => {
    return string.charAt(0).toUpperCase() + string.slice(1);
};

export const getRandomColor = () => {
    const letters = '0123456789ABCDEF';
    let color = '#';

    const randomValues = new Uint8Array(6);
    window.crypto.getRandomValues(randomValues);

    const highIndex = Math.floor(randomValues[0] % 3) * 2;
    for (let i = 0; i < 6; i++) {
        let random;
        if (i >= highIndex && i < highIndex + 2) {
            random = 8 + Math.floor(randomValues[i] % 8);
        } else {
            random = randomValues[i] % 16;
        }
        color += letters[random];
    }
    return color;
};

export function generateUniqueKey(prefix = 'unique-key-') {
    const size = 16;
    const randomValues = window.crypto.getRandomValues(new Uint8Array(size));
    const randomHex = Array.from(randomValues).map(b => b.toString(16).padStart(2, '0')).join('');
    return `${prefix}${randomHex}`;
}
