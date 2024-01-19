export const capitalizeFirstLetter = (string: string) => {
    return string.charAt(0).toUpperCase() + string.slice(1);
};

export const getRandomColor = () => {
    const letters = '0123456789ABCDEF';
    let color = '#';
    for (let i = 0; i < 6; i++) {
        const random = window.crypto.getRandomValues(new Uint8Array(1))[0];
        color += letters[random % 16];
    }
    return color;
};


export function generateUniqueKey(prefix = 'unique-key-') {
    const random = window.crypto.getRandomValues(new Uint8Array(1))[0];
    return `${prefix}${random}`;
}