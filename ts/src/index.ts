import crypto from 'crypto';

interface BlockShape {
    hash: string;
    prevHash: string;
    height: number;
    data: string;
}

class Block implements BlockShape {
    public hash: string;
    constructor(
        public prevHash: string,
        public height: number,
        public data: string
    ) {
        this.hash = Block.calculateHash(prevHash, height, data);
    }

    static calculateHash(prevHash: string, height: number, data: string): string {
        const toHash = `${prevHash}${height}${data}`;
        const hash = crypto.createHash('sha256'); 
        return hash.update(toHash).digest('hex');
    }
}

class Blockchain {
    private blocks: Block[];
    constructor() {
        this.blocks = [];
    }

    private getPrevHash() {
        if (this.blocks.length === 0) return "";
        return this.blocks[this.blocks.length - 1].hash;
    }

    public addBlock(data: string) {
        const block = new Block(this.getPrevHash(), this.blocks.length+1, data);
        this.blocks.push(block);
    }

    public getBlocks() {
        return this.blocks;
    }
}

const blockchain = new Blockchain();

blockchain.addBlock("First One");
blockchain.addBlock("Second One");
blockchain.addBlock("Third One");

// getBlocks() is public, so we can modify private blocks array
blockchain.getBlocks().push(new Block("xxxxxxxx", 11111, "HACKED"));

console.log(blockchain.getBlocks());